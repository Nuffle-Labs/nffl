package aggregator

import (
	"context"
	"sync"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

type OperatorRegistrationsService interface {
	operatorsinfo.OperatorsInfoService

	GetOperatorInfoById(ctx context.Context, operatorId types.OperatorId) (operatorInfo types.OperatorInfo, operatorFound bool)
}

type OperatorRegistrationsServiceInMemory struct {
	avsRegistrySubscriber avsregistry.AvsRegistrySubscriber
	avsRegistryReader     avsregistry.AvsRegistryReader
	logger                logging.Logger
	queryByAddrC          chan<- queryByAddr
	queryByIdC            chan<- queryById

	idToAddr    map[types.OperatorId]common.Address
	addrToId    map[common.Address]types.OperatorId
	pubkeysById map[types.OperatorId]types.OperatorPubkeys
	socketById  map[types.OperatorId]types.Socket
}

type queryByAddr struct {
	operatorAddr common.Address
	respC        chan<- resp
}

type queryById struct {
	operatorId types.OperatorId
	respC      chan<- resp
}

type resp struct {
	operatorInfo   types.OperatorInfo
	operatorExists bool
}

var _ OperatorRegistrationsService = (*OperatorRegistrationsServiceInMemory)(nil)

func NewOperatorRegistrationsServiceInMemory(
	ctx context.Context,
	avsRegistrySubscriber avsregistry.AvsRegistrySubscriber,
	avsRegistryReader avsregistry.AvsRegistryReader,
	logger logging.Logger,
) (*OperatorRegistrationsServiceInMemory, error) {
	queryByAddrC := make(chan queryByAddr)
	queryByIdC := make(chan queryById)

	ors := &OperatorRegistrationsServiceInMemory{
		avsRegistrySubscriber: avsRegistrySubscriber,
		avsRegistryReader:     avsRegistryReader,
		logger:                logger,
		queryByAddrC:          queryByAddrC,
		queryByIdC:            queryByIdC,
		idToAddr:              make(map[types.OperatorId]common.Address),
		addrToId:              make(map[common.Address]types.OperatorId),
		pubkeysById:           make(map[types.OperatorId]types.OperatorPubkeys),
		socketById:            make(map[types.OperatorId]types.Socket),
	}
	err := ors.asyncInit(ctx, queryByAddrC, queryByIdC)
	if err != nil {
		return nil, err
	}

	return ors, nil
}

func (ors *OperatorRegistrationsServiceInMemory) asyncInit(ctx context.Context, queryByAddrC chan queryByAddr, queryByIdC chan queryById) error {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	initErrC := make(chan error)
	ors.startServiceInGoroutine(ctx, queryByAddrC, queryByIdC, &wg, initErrC)

	return <-initErrC
}

func (ors *OperatorRegistrationsServiceInMemory) startServiceInGoroutine(ctx context.Context, queryByAddrC <-chan queryByAddr, queryByIdC <-chan queryById, wg *sync.WaitGroup, initErrC chan<- error) {
	wg.Add(1)

	go func() {
		ors.logger.Debug("Subscribing to new pubkey and socket registration events", "service", "OperatorRegistrationsServiceInMemory")
		newPubkeyRegistrationC, newPubkeyRegistrationSub, err := ors.avsRegistrySubscriber.SubscribeToNewPubkeyRegistrations()
		if err != nil {
			ors.logger.Error("Error opening websocket subscription for new pubkey registrations", "err", err, "service", "OperatorRegistrationsServiceInMemory")
			wg.Done()
			initErrC <- err
			return
		}

		newSocketRegistrationC, newSocketRegistrationSub, err := ors.avsRegistrySubscriber.SubscribeToOperatorSocketUpdates()
		if err != nil {
			ors.logger.Error("Error opening websocket subscription for new socket registrations", "err", err, "service", "OperatorRegistrationsServiceInMemory")
			wg.Done()
			initErrC <- err
			return
		}

		err = ors.queryPastRegisteredOperators(ctx)
		if err != nil {
			wg.Done()
			initErrC <- err
			return
		}
		wg.Done()
		close(initErrC)

		for {
			select {
			case <-ctx.Done():
				ors.logger.Info("OperatorRegistrationsServiceInMemory: Context cancelled, exiting")
				return

			case err := <-newPubkeyRegistrationSub.Err():
				newPubkeyRegistrationSub.Unsubscribe()
				ors.logger.Error("Error in websocket subscription for new pubkey registration events", "err", err, "service", "OperatorRegistrationsServiceInMemory")

			case err := <-newSocketRegistrationSub.Err():
				newSocketRegistrationSub.Unsubscribe()
				ors.logger.Error("Error in websocket subscription for new socket registration events", "err", err, "service", "OperatorRegistrationsServiceInMemory")

			case newPubkeyRegistrationEvent := <-newPubkeyRegistrationC:
				pubkeys := types.OperatorPubkeys{
					G1Pubkey: bls.NewG1Point(newPubkeyRegistrationEvent.PubkeyG1.X, newPubkeyRegistrationEvent.PubkeyG1.Y),
					G2Pubkey: bls.NewG2Point(newPubkeyRegistrationEvent.PubkeyG2.X, newPubkeyRegistrationEvent.PubkeyG2.Y),
				}
				operatorId := types.OperatorIdFromG1Pubkey(pubkeys.G1Pubkey)
				operatorAddr := newPubkeyRegistrationEvent.Operator

				ors.idToAddr[operatorId] = operatorAddr
				ors.addrToId[operatorAddr] = operatorId
				ors.pubkeysById[operatorId] = pubkeys

				ors.logger.Debug("Added operator info to dict",
					"service", "OperatorRegistrationsServiceInMemory",
					"block", newPubkeyRegistrationEvent.Raw.BlockNumber,
					"operatorAddr", operatorAddr,
					"operatorId", operatorId,
					"G1pubkey", pubkeys.G1Pubkey,
					"G2pubkey", pubkeys.G2Pubkey,
				)

			case newSocketRegistrationEvent := <-newSocketRegistrationC:
				operatorId := types.OperatorId(newSocketRegistrationEvent.OperatorId)
				socket := types.Socket(newSocketRegistrationEvent.Socket)
				ors.logger.Debug("Received new socket registration event", "service", "OperatorRegistrationsServiceInMemory", "operatorId", operatorId, "socket", socket)

				ors.socketById[operatorId] = socket

			case q := <-queryByAddrC:
				operatorId, idExists := ors.addrToId[q.operatorAddr]
				pubkeys, pubkeysExist := ors.pubkeysById[operatorId]
				socket, socketExists := ors.socketById[operatorId]

				operatorInfo := types.OperatorInfo{
					Pubkeys: pubkeys,
					Socket:  socket,
				}
				q.respC <- resp{operatorInfo, idExists && pubkeysExist && socketExists}

			case q := <-queryByIdC:
				pubkeys, pubkeysExist := ors.pubkeysById[q.operatorId]
				socket, socketExists := ors.socketById[q.operatorId]

				operatorInfo := types.OperatorInfo{
					Pubkeys: pubkeys,
					Socket:  socket,
				}
				q.respC <- resp{operatorInfo, pubkeysExist && socketExists}
			}
		}
	}()
}

func (ors *OperatorRegistrationsServiceInMemory) queryPastRegisteredOperators(ctx context.Context) error {
	alreadyRegisteredOperatorAddrs, alreadyRegisteredOperatorPubkeys, err := ors.avsRegistryReader.QueryExistingRegisteredOperatorPubKeys(ctx, nil, nil)
	if err != nil {
		ors.logger.Error("Error querying existing registered operators", "err", err, "service", "OperatorRegistrationsServiceInMemory")
		return err
	}

	socketById, err := ors.avsRegistryReader.QueryExistingRegisteredOperatorSockets(ctx, nil, nil)
	if err != nil {
		ors.logger.Error("Error querying existing registered operator sockets", "err", err, "service", "OperatorRegistrationsServiceInMemory")
		return err
	}

	for i, operatorAddr := range alreadyRegisteredOperatorAddrs {
		operatorPubkeys := alreadyRegisteredOperatorPubkeys[i]
		operatorId := types.OperatorIdFromG1Pubkey(operatorPubkeys.G1Pubkey)

		ors.idToAddr[operatorId] = operatorAddr
		ors.addrToId[operatorAddr] = operatorId
		ors.pubkeysById[operatorId] = operatorPubkeys
		ors.socketById[operatorId] = socketById[operatorId]
	}

	return nil
}

func (ors *OperatorRegistrationsServiceInMemory) GetOperatorInfo(ctx context.Context, operator common.Address) (types.OperatorInfo, bool) {
	respC := make(chan resp)
	ors.queryByAddrC <- queryByAddr{operator, respC}

	select {
	case <-ctx.Done():
		return types.OperatorInfo{}, false
	case resp := <-respC:
		return resp.operatorInfo, resp.operatorExists
	}
}

func (ors *OperatorRegistrationsServiceInMemory) GetOperatorInfoById(ctx context.Context, operatorId types.OperatorId) (types.OperatorInfo, bool) {
	respC := make(chan resp)
	ors.queryByIdC <- queryById{operatorId, respC}

	select {
	case <-ctx.Done():
		return types.OperatorInfo{}, false
	case resp := <-respC:
		return resp.operatorInfo, resp.operatorExists
	}
}
