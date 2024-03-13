package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/NethermindEth/near-sffl/aggregator/types"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	"github.com/NethermindEth/near-sffl/core/database/models"
)

type Databaser interface {
	Close() error
	StoreStateRootUpdate(stateRootUpdateMessage servicemanager.StateRootUpdateMessage) error
	FetchStateRootUpdate(rollupId uint32, blockHeight uint64, stateRootUpdateMessage *servicemanager.StateRootUpdateMessage) error
	StoreStateRootUpdateAggregation(stateRootUpdateMessage servicemanager.StateRootUpdateMessage, aggregation types.MessageBlsAggregationServiceResponse) error
	FetchStateRootUpdateAggregation(rollupId uint32, blockHeight uint64, aggregation *types.MessageBlsAggregationServiceResponse) error
	StoreOperatorSetUpdate(operatorSetUpdateMessage registryrollup.OperatorSetUpdateMessage) error
	FetchOperatorSetUpdate(id uint64, operatorSetUpdateMessage *registryrollup.OperatorSetUpdateMessage) error
	StoreOperatorSetUpdateAggregation(operatorSetUpdateMessage registryrollup.OperatorSetUpdateMessage, aggregation types.MessageBlsAggregationServiceResponse) error
	FetchOperatorSetUpdateAggregation(id uint64, aggregation *types.MessageBlsAggregationServiceResponse) error
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
	tx := d.db.Create(&models.StateRootUpdateMessage{
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

func (d *Database) StoreStateRootUpdateAggregation(stateRootUpdateMessage servicemanager.StateRootUpdateMessage, aggregation types.MessageBlsAggregationServiceResponse) error {
	err := d.db.
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

func (d *Database) FetchStateRootUpdateAggregation(rollupId uint32, blockHeight uint64, aggregation *types.MessageBlsAggregationServiceResponse) error {
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
	tx := d.db.Create(&models.OperatorSetUpdateMessage{
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

func (d *Database) StoreOperatorSetUpdateAggregation(operatorSetUpdateMessage registryrollup.OperatorSetUpdateMessage, aggregation types.MessageBlsAggregationServiceResponse) error {
	err := d.db.
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

func (d *Database) FetchOperatorSetUpdateAggregation(id uint64, aggregation *types.MessageBlsAggregationServiceResponse) error {
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
