package aggregator

import (
	"context"
	"sync"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/operatorpubkeys"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

type OperatorRegistrationsService interface {
	operatorpubkeys.OperatorPubkeysService

	GetOperatorPubkeysById(ctx context.Context, operatorId types.OperatorId) (operatorPubkeys types.OperatorPubkeys, operatorFound bool)
}

type OperatorRegistrationsServiceInMemory struct {
	avsRegistrySubscriber avsregistry.AvsRegistrySubscriber
	avsRegistryReader     avsregistry.AvsRegistryReader
	logger                logging.Logger
	queryByAddrC          chan<- queryByAddr
	queryByIdC            chan<- queryById
}

type queryByAddr struct {
	operatorAddr common.Address
	// channel through which to receive the response (operator pubkeys)
	respC chan<- resp
}
type queryById struct {
	operatorId types.OperatorId
	// channel through which to receive the response (operator pubkeys)
	respC chan<- resp
}

type resp struct {
	operatorPubkeys types.OperatorPubkeys
	// false if operators were not present in the pubkey dict
	operatorExists bool
}

var _ operatorpubkeys.OperatorPubkeysService = (*OperatorRegistrationsServiceInMemory)(nil)

// NewOperatorRegistrationsServiceInMemory constructs a OperatorRegistrationsServiceInMemory and starts it in a goroutine.
// It takes a context as argument because the "backfilling" of the database is done inside this constructor,
// so we wait for all past NewPubkeyRegistration events to be queried and the db to be filled before returning the service.
// The constructor is thus following a RAII-like pattern, of initializing the serving during construction.
// Using a separate initialize() function might lead to some users forgetting to call it and the service not behaving properly.
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
	}
	err := ors.asyncInit(ctx, queryByAddrC, queryByIdC)
	if err != nil {
		return nil, err
	}

	return ors, nil
}

// asyncInit parks caller & schedules initialization of the inmemory pubkey dict,
// which requires querying the past events of the pubkey registration contract
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
		ors.logger.Debug("Subscribing to new pubkey registration events on blsApkRegistry contract", "service", "OperatorRegistrationsServiceInMemory")
		newPubkeyRegistrationC, newPubkeyRegistrationSub, err := ors.avsRegistrySubscriber.SubscribeToNewPubkeyRegistrations()
		if err != nil {
			ors.logger.Error("Error opening websocket subscription for new pubkey registrations", "err", err, "service", "OperatorRegistrationsServiceInMemory")
			wg.Done()
			initErrC <- err
			return
		}

		pubkeyByAddrDict, pubkeyByIdDict, err := ors.queryPastRegisteredOperators(ctx)
		if err != nil {
			wg.Done()
			initErrC <- err
			return
		}
		// The constructor can return after we have backfilled the db by querying the events of operators that have registered with the blsApkRegistry
		// before the block at which we started the ws subscription above
		wg.Done()
		close(initErrC)

		for {
			select {
			case <-ctx.Done():
				ors.logger.Info("OperatorRegistrationsServiceInMemory: Context cancelled, exiting")
				return

			// This shall not really happen unless ctx was canceled & this came first.
			case err := <-newPubkeyRegistrationSub.Err():
				// Just report
				newPubkeyRegistrationSub.Unsubscribe()
				ors.logger.Error("Error in safe client subscription", "err", err, "service", "OperatorRegistrationsServiceInMemory")

			case newPubkeyRegistrationEvent := <-newPubkeyRegistrationC:
				pubkeys := types.OperatorPubkeys{
					G1Pubkey: bls.NewG1Point(newPubkeyRegistrationEvent.PubkeyG1.X, newPubkeyRegistrationEvent.PubkeyG1.Y),
					G2Pubkey: bls.NewG2Point(newPubkeyRegistrationEvent.PubkeyG2.X, newPubkeyRegistrationEvent.PubkeyG2.Y),
				}
				operatorId := types.OperatorIdFromPubkey(pubkeys.G1Pubkey)
				operatorAddr := newPubkeyRegistrationEvent.Operator

				pubkeyByAddrDict[operatorAddr] = pubkeys
				pubkeyByIdDict[operatorId] = pubkeys

				ors.logger.Debug("Added operator pubkeys to pubkey dict",
					"service", "OperatorRegistrationsServiceInMemory",
					"block", newPubkeyRegistrationEvent.Raw.BlockNumber,
					"operatorAddr", operatorAddr,
					"operatorId", operatorId,
					"G1pubkey", pubkeyByAddrDict[operatorAddr].G1Pubkey,
					"G2pubkey", pubkeyByAddrDict[operatorAddr].G2Pubkey,
				)

			// Receive a queryByAddr from GetOperatorPubkeys
			case operatorPubkeyQuery := <-queryByAddrC:
				pubkeys, ok := pubkeyByAddrDict[operatorPubkeyQuery.operatorAddr]
				operatorPubkeyQuery.respC <- resp{pubkeys, ok}

			// Receive a queryById from GetOperatorPubkeysById
			case operatorPubkeyQuery := <-queryByIdC:
				pubkeys, ok := pubkeyByIdDict[operatorPubkeyQuery.operatorId]
				operatorPubkeyQuery.respC <- resp{pubkeys, ok}
			}
		}
	}()
}

func (ors *OperatorRegistrationsServiceInMemory) queryPastRegisteredOperators(ctx context.Context) (map[common.Address]types.OperatorPubkeys, map[types.OperatorId]types.OperatorPubkeys, error) {
	// Querying with nil startBlock and stopBlock will return all events. It doesn't matter if we queryByAddr some events that we will receive again in the websocket,
	// since we will just overwrite the pubkey dict with the same values.
	alreadyRegisteredOperatorAddrs, alreadyRegisteredOperatorPubkeys, err := ors.avsRegistryReader.QueryExistingRegisteredOperatorPubKeys(ctx, nil, nil)
	if err != nil {
		ors.logger.Error("Error querying existing registered operators", "err", err, "service", "OperatorRegistrationsServiceInMemory")
		return nil, nil, err
	}

	ors.logger.Debug("List of queried operator registration events in blsApkRegistry", "alreadyRegisteredOperatorAddr", alreadyRegisteredOperatorAddrs, "service", "OperatorRegistrationsServiceInMemory")

	pubkeyByAddrDict := make(map[common.Address]types.OperatorPubkeys)
	pubkeyByIdDict := make(map[types.OperatorId]types.OperatorPubkeys)
	for i, operatorAddr := range alreadyRegisteredOperatorAddrs {
		operatorPubkeys := alreadyRegisteredOperatorPubkeys[i]
		pubkeyByAddrDict[operatorAddr] = operatorPubkeys

		operatorId := types.OperatorIdFromPubkey(operatorPubkeys.G1Pubkey)
		pubkeyByIdDict[operatorId] = operatorPubkeys
	}

	return pubkeyByAddrDict, pubkeyByIdDict, nil
}

func (ors *OperatorRegistrationsServiceInMemory) GetOperatorPubkeys(ctx context.Context, operator common.Address) (types.OperatorPubkeys, bool) {
	respC := make(chan resp)
	ors.queryByAddrC <- queryByAddr{operator, respC}

	select {
	case <-ctx.Done():
		return types.OperatorPubkeys{}, false
	case resp := <-respC:
		return resp.operatorPubkeys, resp.operatorExists
	}
}

func (ors *OperatorRegistrationsServiceInMemory) GetOperatorPubkeysById(ctx context.Context, operatorId types.OperatorId) (types.OperatorPubkeys, bool) {
	respC := make(chan resp)
	ors.queryByIdC <- queryById{operatorId, respC}

	select {
	case <-ctx.Done():
		return types.OperatorPubkeys{}, false
	case resp := <-respC:
		return resp.operatorPubkeys, resp.operatorExists
	}
}
