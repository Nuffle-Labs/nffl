package database

import (
	"errors"
	"math"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/NethermindEth/near-sffl/aggregator/database/models"
	aggtypes "github.com/NethermindEth/near-sffl/aggregator/types"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
)

type Databaser interface {
	Close() error
	StoreStateRootUpdate(stateRootUpdateMessage servicemanager.StateRootUpdateMessage) error
	FetchStateRootUpdate(rollupId uint32, blockHeight uint64, stateRootUpdateMessage *servicemanager.StateRootUpdateMessage) error
	StoreStateRootUpdateAggregation(stateRootUpdateMessage servicemanager.StateRootUpdateMessage, aggregation aggtypes.MessageBlsAggregationServiceResponse) error
	FetchStateRootUpdateAggregation(rollupId uint32, blockHeight uint64, aggregation *aggtypes.MessageBlsAggregationServiceResponse) error
	StoreOperatorSetUpdate(operatorSetUpdateMessage registryrollup.OperatorSetUpdateMessage) error
	FetchOperatorSetUpdate(id uint64, operatorSetUpdateMessage *registryrollup.OperatorSetUpdateMessage) error
	StoreOperatorSetUpdateAggregation(operatorSetUpdateMessage registryrollup.OperatorSetUpdateMessage, aggregation aggtypes.MessageBlsAggregationServiceResponse) error
	FetchOperatorSetUpdateAggregation(id uint64, aggregation *aggtypes.MessageBlsAggregationServiceResponse) error
	FetchCheckpointMessages(fromTimestamp uint64, toTimestamp uint64, result *coretypes.CheckpointMessages) error
}

type Database struct {
	db     *gorm.DB
	dbPath string
}

func NewDatabase(dbPath string) (*Database, error) {
	if dbPath == "" {
		dbPath = ":memory:"
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&models.MessageBlsAggregation{},
		&models.StateRootUpdateMessage{},
		&models.OperatorSetUpdateMessage{},
	)

	return &Database{
		db:     db,
		dbPath: dbPath,
	}, nil
}

func (d *Database) Close() error {
	db, err := d.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (d *Database) StoreStateRootUpdate(stateRootUpdateMessage servicemanager.StateRootUpdateMessage) error {
	tx := d.db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&models.StateRootUpdateMessage{
			RollupId:    stateRootUpdateMessage.RollupId,
			BlockHeight: stateRootUpdateMessage.BlockHeight,
			Timestamp:   stateRootUpdateMessage.Timestamp,
			StateRoot:   stateRootUpdateMessage.StateRoot[:],
		})

	return tx.Error
}

func (d *Database) FetchStateRootUpdate(rollupId uint32, blockHeight uint64, stateRootUpdateMessage *servicemanager.StateRootUpdateMessage) error {
	var model models.StateRootUpdateMessage

	tx := d.db.
		Where("rollup_id = ?", rollupId).
		Where("block_height = ?", blockHeight).
		First(&model)
	if tx.Error != nil {
		return tx.Error
	}

	*stateRootUpdateMessage = model.Parse()

	return nil
}

func (d *Database) StoreStateRootUpdateAggregation(stateRootUpdateMessage servicemanager.StateRootUpdateMessage, aggregation aggtypes.MessageBlsAggregationServiceResponse) error {
	err := d.db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Model(&models.StateRootUpdateMessage{}).
		Where("rollup_id = ?", stateRootUpdateMessage.RollupId).
		Where("block_height = ?", stateRootUpdateMessage.BlockHeight).
		Association("Aggregation").
		Replace(&models.MessageBlsAggregation{
			EthBlockNumber:               aggregation.EthBlockNumber,
			MessageDigest:                aggregation.MessageDigest[:],
			NonSignersPubkeysG1:          aggregation.NonSignersPubkeysG1,
			QuorumApksG1:                 aggregation.QuorumApksG1,
			SignersApkG2:                 aggregation.SignersApkG2,
			SignersAggSigG1:              aggregation.SignersAggSigG1,
			NonSignerQuorumBitmapIndices: aggregation.NonSignerQuorumBitmapIndices,
			QuorumApkIndices:             aggregation.QuorumApkIndices,
			TotalStakeIndices:            aggregation.TotalStakeIndices,
			NonSignerStakeIndices:        aggregation.NonSignerStakeIndices,
		})
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) FetchStateRootUpdateAggregation(rollupId uint32, blockHeight uint64, aggregation *aggtypes.MessageBlsAggregationServiceResponse) error {
	var model models.StateRootUpdateMessage

	tx := d.db.
		Preload("Aggregation").
		Model(&model).
		Where("rollup_id = ?", rollupId).
		Where("block_height = ?", blockHeight).
		First(&model)
	if tx.Error != nil {
		return tx.Error
	}

	*aggregation = model.Aggregation.Parse()

	return nil
}

func (d *Database) StoreOperatorSetUpdate(operatorSetUpdateMessage registryrollup.OperatorSetUpdateMessage) error {
	tx := d.db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&models.OperatorSetUpdateMessage{
			UpdateId:  operatorSetUpdateMessage.Id,
			Timestamp: operatorSetUpdateMessage.Timestamp,
			Operators: operatorSetUpdateMessage.Operators,
		})

	return tx.Error
}

func (d *Database) FetchOperatorSetUpdate(id uint64, operatorSetUpdateMessage *registryrollup.OperatorSetUpdateMessage) error {
	var model models.OperatorSetUpdateMessage

	tx := d.db.
		Where("update_id = ?", id).
		First(&model)
	if tx.Error != nil {
		return tx.Error
	}
	*operatorSetUpdateMessage = model.Parse()

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (d *Database) StoreOperatorSetUpdateAggregation(operatorSetUpdateMessage registryrollup.OperatorSetUpdateMessage, aggregation aggtypes.MessageBlsAggregationServiceResponse) error {
	err := d.db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Model(&models.OperatorSetUpdateMessage{}).
		Where("update_id = ?", operatorSetUpdateMessage.Id).
		Association("Aggregation").
		Replace(&models.MessageBlsAggregation{
			EthBlockNumber:               aggregation.EthBlockNumber,
			MessageDigest:                aggregation.MessageDigest[:],
			NonSignersPubkeysG1:          aggregation.NonSignersPubkeysG1,
			QuorumApksG1:                 aggregation.QuorumApksG1,
			SignersApkG2:                 aggregation.SignersApkG2,
			SignersAggSigG1:              aggregation.SignersAggSigG1,
			NonSignerQuorumBitmapIndices: aggregation.NonSignerQuorumBitmapIndices,
			QuorumApkIndices:             aggregation.QuorumApkIndices,
			TotalStakeIndices:            aggregation.TotalStakeIndices,
			NonSignerStakeIndices:        aggregation.NonSignerStakeIndices,
		})
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) FetchOperatorSetUpdateAggregation(id uint64, aggregation *aggtypes.MessageBlsAggregationServiceResponse) error {
	var model models.OperatorSetUpdateMessage

	tx := d.db.
		Preload("Aggregation").
		Model(&model).
		Where("update_id = ?", id).
		First(&model)
	if tx.Error != nil {
		return tx.Error
	}

	*aggregation = model.Aggregation.Parse()

	return nil
}

func (d *Database) FetchCheckpointMessages(fromTimestamp uint64, toTimestamp uint64, result *coretypes.CheckpointMessages) error {
	if fromTimestamp > math.MaxInt64 || toTimestamp > math.MaxInt64 {
		return errors.New("timestamp does not fit in int64")
	}

	var stateRootUpdates []models.StateRootUpdateMessage

	tx := d.db.
		Preload("Aggregation").
		Model(&models.StateRootUpdateMessage{}).
		Where("timestamp >= ?", fromTimestamp).
		Where("timestamp <= ?", toTimestamp).
		First(&stateRootUpdates)
	if tx.Error != nil {
		return tx.Error
	}

	var operatorSetUpdates []models.OperatorSetUpdateMessage

	tx = d.db.
		Preload("Aggregation").
		Model(&models.OperatorSetUpdateMessage{}).
		Where("timestamp >= ?", fromTimestamp).
		Where("timestamp <= ?", toTimestamp).
		First(&operatorSetUpdates)
	if tx.Error != nil {
		return tx.Error
	}

	stateRootUpdateMessages := make([]servicemanager.StateRootUpdateMessage, 0, len(stateRootUpdates))
	stateRootUpdateMessageAggregations := make([]aggtypes.MessageBlsAggregationServiceResponse, 0, len(stateRootUpdates))
	operatorSetUpdateMessages := make([]registryrollup.OperatorSetUpdateMessage, 0, len(operatorSetUpdates))
	operatorSetUpdateMessageAggregations := make([]aggtypes.MessageBlsAggregationServiceResponse, 0, len(operatorSetUpdates))

	for _, stateRootUpdate := range stateRootUpdates {
		agg := stateRootUpdate.Aggregation

		stateRootUpdateMessages = append(stateRootUpdateMessages, stateRootUpdate.Parse())
		stateRootUpdateMessageAggregations = append(stateRootUpdateMessageAggregations, agg.Parse())
	}

	for _, operatorSetUpdate := range operatorSetUpdates {
		agg := operatorSetUpdate.Aggregation

		operatorSetUpdateMessages = append(operatorSetUpdateMessages, operatorSetUpdate.Parse())
		operatorSetUpdateMessageAggregations = append(operatorSetUpdateMessageAggregations, agg.Parse())
	}

	*result = coretypes.CheckpointMessages{
		StateRootUpdateMessages:              stateRootUpdateMessages,
		StateRootUpdateMessageAggregations:   stateRootUpdateMessageAggregations,
		OperatorSetUpdateMessages:            operatorSetUpdateMessages,
		OperatorSetUpdateMessageAggregations: operatorSetUpdateMessageAggregations,
	}

	return nil
}
