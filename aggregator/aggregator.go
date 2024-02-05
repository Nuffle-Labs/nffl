package aggregator

import (
	"context"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	oppubkeysserv "github.com/Layr-Labs/eigensdk-go/services/operatorpubkeys"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	badger "github.com/dgraph-io/badger/v4"

	"github.com/NethermindEth/near-sffl/aggregator/types"
	"github.com/NethermindEth/near-sffl/core"
	"github.com/NethermindEth/near-sffl/core/chainio"
	"github.com/NethermindEth/near-sffl/core/config"

	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
)

const (
	// number of blocks after which a task is considered expired
	// this hardcoded here because it's also hardcoded in the contracts, but should
	// ideally be fetched from the contracts
	taskChallengeWindowBlock = 100
	blockTimeSeconds         = 12 * time.Second
	avsName                  = "super-fast-finality-layer"
)

// Aggregator sends checkpoint tasks onchain, then listens for operator signed TaskResponses.
// It aggregates responses signatures, and if any of the TaskResponses reaches the QuorumThreshold for each quorum
// (currently we only use a single quorum of the ERC20Mock token), it sends the aggregated TaskResponse and signature onchain.
//
// The signature is checked in the BLSSignatureChecker.sol contract, which expects a
//
//	struct NonSignerStakesAndSignature {
//		uint32[] nonSignerQuorumBitmapIndices;
//		BN254.G1Point[] nonSignerPubkeys;
//		BN254.G1Point[] quorumApks;
//		BN254.G2Point apkG2;
//		BN254.G1Point sigma;
//		uint32[] quorumApkIndices;
//		uint32[] totalStakeIndices;
//		uint32[][] nonSignerStakeIndices; // nonSignerStakeIndices[quorumNumberIndex][nonSignerIndex]
//	}
//
// A task can only be responded onchain by having enough operators sign on it such that their stake in each quorum reaches the QuorumThreshold.
// In order to verify this onchain, the Registry contracts store the history of stakes and aggregate pubkeys (apks) for each operators and each quorum. These are
// updated everytime an operator registers or deregisters with the BLSRegistryCoordinatorWithIndices.sol contract, or calls UpdateStakes() on the StakeRegistry.sol contract,
// after having received new delegated shares or having delegated shares removed by stakers queuing withdrawals. Each of these pushes to their respective datatype array a new entry.
//
// This is true for quorumBitmaps (represent the quorums each operator is opted into), quorumApks (apks per quorum), totalStakes (total stake per quorum), and nonSignerStakes (stake per quorum per operator).
// The 4 "indices" in NonSignerStakesAndSignature basically represent the index at which to fetch their respective data, given a blockNumber at which the task was created.
// Note that different data types might have different indices, since for eg QuorumBitmaps are updated for operators registering/deregistering, but not for UpdateStakes.
// Thankfully, we have deployed a helper contract BLSOperatorStateRetriever.sol whose function getCheckSignaturesIndices() can be used to fetch the indices given a block number.
//
// The 4 other fields nonSignerPubkeys, quorumApks, apkG2, and sigma, however, must be computed individually.
// apkG2 and sigma are just the aggregated signature and pubkeys of the operators who signed the task response (aggregated over all quorums, so individual signatures might be duplicated).
// quorumApks are the G1 aggregated pubkeys of the operators who signed the task response, but one per quorum, as opposed to apkG2 which is summed over all quorums.
// nonSignerPubkeys are the G1 pubkeys of the operators who did not sign the task response, but were opted into the quorum at the blocknumber at which the task was created.
// Upon sending a task onchain (or receiving a CheckpointTaskCreated Event if the tasks were sent by an external task generator), the aggregator can get the list of all operators opted into each quorum at that
// block number by calling the getOperatorState() function of the BLSOperatorStateRetriever.sol contract.
type Aggregator struct {
	logger               logging.Logger
	serverIpPortAddr     string
	restServerIpPortAddr string
	databasePath         string
	avsWriter            chainio.AvsWriterer
	// aggregation related fields
	taskBlsAggregationService    blsagg.BlsAggregationService
	messageBlsAggregationService MessageBlsAggregationService
	tasks                        map[types.TaskIndex]taskmanager.CheckpointTask
	tasksMu                      sync.RWMutex
	taskResponses                map[types.TaskIndex]map[sdktypes.TaskResponseDigest]taskmanager.CheckpointTaskResponse
	taskResponsesMu              sync.RWMutex
	database                     *badger.DB
	databaseMu                   sync.RWMutex
	stateRootUpdates             map[types.MessageDigest]servicemanager.StateRootUpdateMessage
	stateRootUpdatesMu           sync.RWMutex
}

// NewAggregator creates a new Aggregator with the provided config.
func NewAggregator(c *config.Config) (*Aggregator, error) {
	avsReader, err := chainio.BuildAvsReaderFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create avsReader", "err", err)
		return nil, err
	}

	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	chainioConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 c.EthHttpRpcUrl,
		EthWsUrl:                   c.EthWsRpcUrl,
		RegistryCoordinatorAddr:    c.SFFLRegistryCoordinatorAddr.String(),
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddr.String(),
		AvsName:                    avsName,
		PromMetricsIpPortAddress:   ":9090",
	}
	clients, err := clients.BuildAll(chainioConfig, c.AggregatorAddress, c.SignerFn, c.Logger)
	if err != nil {
		c.Logger.Errorf("Cannot create sdk clients", "err", err)
		return nil, err
	}

	operatorPubkeysService := oppubkeysserv.NewOperatorPubkeysServiceInMemory(context.Background(), clients.AvsRegistryChainSubscriber, clients.AvsRegistryChainReader, c.Logger)
	avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(avsReader, operatorPubkeysService, c.Logger)
	taskBlsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, c.Logger)
	messageBlsAggregationService := NewMessageBlsAggregatorService(avsRegistryService, clients.EthHttpClient, c.Logger)

	return &Aggregator{
		logger:                       c.Logger,
		serverIpPortAddr:             c.AggregatorServerIpPortAddr,
		restServerIpPortAddr:         c.AggregatorRestServerIpPortAddr,
		databasePath:                 c.AggregatorDatabasePath,
		avsWriter:                    avsWriter,
		taskBlsAggregationService:    taskBlsAggregationService,
		messageBlsAggregationService: messageBlsAggregationService,
		tasks:                        make(map[types.TaskIndex]taskmanager.CheckpointTask),
		taskResponses:                make(map[types.TaskIndex]map[sdktypes.TaskResponseDigest]taskmanager.CheckpointTaskResponse),
		stateRootUpdates:             make(map[types.MessageDigest]servicemanager.StateRootUpdateMessage),
	}, nil
}

func (agg *Aggregator) Start(ctx context.Context) error {
	agg.logger.Infof("Starting aggregator.")

	agg.logger.Infof("Starting aggregator rpc server.")
	go agg.startServer(ctx)

	agg.logger.Infof("Starting aggregator REST API.")
	go agg.startRestServer(ctx)

	// TODO(soubhik): refactor task generation/sending into a separate function that we can run as goroutine
	ticker := time.NewTicker(10 * time.Second)
	agg.logger.Infof("Aggregator set to send new task every 10 seconds...")
	defer ticker.Stop()

	// ticker doesn't tick immediately, so we send the first task here
	// see https://github.com/golang/go/issues/17601

	// TODO: make this based on the NEAR block
	block := uint64(0)
	_ = agg.sendNewCheckpointTask(block, block)
	block++

	for {
		select {
		case <-ctx.Done():
			return nil
		case blsAggServiceResp := <-agg.taskBlsAggregationService.GetResponseChannel():
			agg.logger.Info("Received response from taskBlsAggregationService", "blsAggServiceResp", blsAggServiceResp)
			agg.sendAggregatedResponseToContract(blsAggServiceResp)
		case blsAggServiceResp := <-agg.messageBlsAggregationService.GetResponseChannel():
			agg.logger.Info("Received response from messageBlsAggregationService", "blsAggServiceResp", blsAggServiceResp)
			agg.handleStateRootUpdateReachedQuorum(blsAggServiceResp)
		case <-ticker.C:
			err := agg.sendNewCheckpointTask(block, block)
			block++
			if err != nil {
				// we log the errors inside sendNewCheckpointTask() so here we just continue to the next task
				continue
			}
		}
	}
}

func (agg *Aggregator) sendAggregatedResponseToContract(blsAggServiceResp blsagg.BlsAggregationServiceResponse) {
	// TODO: check if blsAggServiceResp contains an err
	if blsAggServiceResp.Err != nil {
		agg.logger.Error("BlsAggregationServiceResponse contains an error", "err", blsAggServiceResp.Err)
		// panicing to help with debugging (fail fast), but we shouldn't panic if we run this in production
		panic(blsAggServiceResp.Err)
	}
	nonSignerPubkeys := []taskmanager.BN254G1Point{}
	for _, nonSignerPubkey := range blsAggServiceResp.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, core.ConvertToBN254G1Point(nonSignerPubkey))
	}
	quorumApks := []taskmanager.BN254G1Point{}
	for _, quorumApk := range blsAggServiceResp.QuorumApksG1 {
		quorumApks = append(quorumApks, core.ConvertToBN254G1Point(quorumApk))
	}
	nonSignerStakesAndSignature := taskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        core.ConvertToBN254G2Point(blsAggServiceResp.SignersApkG2),
		Sigma:                        core.ConvertToBN254G1Point(blsAggServiceResp.SignersAggSigG1.G1Point),
		NonSignerQuorumBitmapIndices: blsAggServiceResp.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             blsAggServiceResp.QuorumApkIndices,
		TotalStakeIndices:            blsAggServiceResp.TotalStakeIndices,
		NonSignerStakeIndices:        blsAggServiceResp.NonSignerStakeIndices,
	}

	agg.logger.Info("Threshold reached. Sending aggregated response onchain.",
		"taskIndex", blsAggServiceResp.TaskIndex,
	)
	agg.tasksMu.RLock()
	task := agg.tasks[blsAggServiceResp.TaskIndex]
	agg.tasksMu.RUnlock()
	agg.taskResponsesMu.RLock()
	taskResponse := agg.taskResponses[blsAggServiceResp.TaskIndex][blsAggServiceResp.TaskResponseDigest]
	agg.taskResponsesMu.RUnlock()
	_, err := agg.avsWriter.SendAggregatedResponse(context.Background(), task, taskResponse, nonSignerStakesAndSignature)
	if err != nil {
		agg.logger.Error("Aggregator failed to respond to task", "err", err)
	}
}

// sendNewCheckpointTask sends a new task to the task manager contract, and updates the Task dict struct
// with the information of operators opted into quorum 0 at the block of task creation.
func (agg *Aggregator) sendNewCheckpointTask(fromTimestamp uint64, toTimestamp uint64) error {
	agg.logger.Info("Aggregator sending new task", "fromTimestamp", fromTimestamp, "toTimestamp", toTimestamp)
	// Send checkpoint to the task manager contract
	newTask, taskIndex, err := agg.avsWriter.SendNewCheckpointTask(context.Background(), fromTimestamp, toTimestamp, types.QUORUM_THRESHOLD_NUMERATOR, types.QUORUM_NUMBERS)
	if err != nil {
		agg.logger.Error("Aggregator failed to send checkpoint", "err", err)
		return err
	}

	agg.tasksMu.Lock()
	agg.tasks[taskIndex] = newTask
	agg.tasksMu.Unlock()

	quorumThresholds := make([]uint32, len(newTask.QuorumNumbers))
	for i, _ := range newTask.QuorumNumbers {
		quorumThresholds[i] = newTask.QuorumThreshold
	}
	// TODO(samlaf): we use seconds for now, but we should ideally pass a blocknumber to the blsAggregationService
	// and it should monitor the chain and only expire the task aggregation once the chain has reached that block number.
	taskTimeToExpiry := taskChallengeWindowBlock * blockTimeSeconds
	agg.taskBlsAggregationService.InitializeNewTask(taskIndex, newTask.TaskCreatedBlock, newTask.QuorumNumbers, quorumThresholds, taskTimeToExpiry)
	return nil
}

func (agg *Aggregator) handleStateRootUpdateReachedQuorum(blsAggServiceResp types.MessageBlsAggregationServiceResponse) {
	agg.stateRootUpdatesMu.RLock()
	msg, ok := agg.stateRootUpdates[blsAggServiceResp.MessageDigest]
	agg.stateRootUpdatesMu.RUnlock()

	defer func() {
		agg.stateRootUpdatesMu.RLock()
		delete(agg.stateRootUpdates, blsAggServiceResp.MessageDigest)
		agg.stateRootUpdatesMu.RUnlock()
	}()

	if !ok {
		agg.logger.Error("Aggregator could not find matching message")
		return
	}

	if blsAggServiceResp.Err != nil {
		agg.logger.Error("Aggregator BLS service returned error", "err", blsAggServiceResp.Err)
		return
	}

	err := agg.storeStateRootUpdate(msg)
	if err != nil {
		agg.logger.Error("Aggregator could not store message")
		return
	}
	agg.storeStateRootUpdateAggregation(msg, blsAggServiceResp)
	if err != nil {
		agg.logger.Error("Aggregator could not store message aggregation")
		return
	}
}
