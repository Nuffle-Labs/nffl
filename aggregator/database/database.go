package database

import (
	"context"
	"errors"
	"log"
	"math"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/NethermindEth/near-sffl/aggregator/database/models"
	"github.com/NethermindEth/near-sffl/core"
	"github.com/NethermindEth/near-sffl/core/types/messages"
)

type Databaser interface {
	core.Metricable

	Close() error
	StoreStateRootUpdate(stateRootUpdateMessage messages.StateRootUpdateMessage) (*models.StateRootUpdateMessage, error)
	FetchStateRootUpdate(rollupId uint32, blockHeight uint64) (*messages.StateRootUpdateMessage, error)
	StoreStateRootUpdateAggregation(stateRootUpdateMessage *models.StateRootUpdateMessage, aggregation messages.MessageBlsAggregation) error
	FetchStateRootUpdateAggregation(rollupId uint32, blockHeight uint64) (*messages.MessageBlsAggregation, error)
	StoreOperatorSetUpdate(operatorSetUpdateMessage messages.OperatorSetUpdateMessage) (*models.OperatorSetUpdateMessage, error)
	FetchOperatorSetUpdate(id uint64) (*messages.OperatorSetUpdateMessage, error)
	StoreOperatorSetUpdateAggregation(operatorSetUpdateMessage *models.OperatorSetUpdateMessage, aggregation messages.MessageBlsAggregation) error
	FetchOperatorSetUpdateAggregation(id uint64) (*messages.MessageBlsAggregation, error)
	FetchCheckpointMessages(fromTimestamp uint64, toTimestamp uint64) (*messages.CheckpointMessages, error)
	DB() *gorm.DB
}

type Database struct {
	db       *gorm.DB
	dbPath   string
	listener EventListener
}

var _ core.Metricable = (*Database)(nil)
var _ Databaser = (*Database)(nil)

func NewDatabase(dbPath string) (*Database, error) {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			Colorful:                  true,
		},
	)

	if dbPath == "" {
		dbPath = "file::memory:?cache=shared"
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{Logger: logger})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.MessageBlsAggregation{},
		&models.StateRootUpdateMessage{},
		&models.OperatorSetUpdateMessage{},
	)
	if err != nil {
		return nil, err
	}

	underlyingDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	underlyingDb.SetMaxOpenConns(1)

	return &Database{
		db:       db,
		dbPath:   dbPath,
		listener: &SelectiveListener{},
	}, nil
}

func (d *Database) Close() error {
	db, err := d.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (d *Database) EnableMetrics(registry *prometheus.Registry) error {
	listener, err := MakeDBMetrics(registry)
	if err != nil {
		return err
	}

	d.listener = listener
	return nil
}

func (d *Database) StoreStateRootUpdate(stateRootUpdateMessage messages.StateRootUpdateMessage) (*models.StateRootUpdateMessage, error) {
	start := time.Now()
	defer func() { d.listener.OnStore(time.Since(start)) }()

	model := models.StateRootUpdateMessage{
		RollupId:            stateRootUpdateMessage.RollupId,
		BlockHeight:         stateRootUpdateMessage.BlockHeight,
		Timestamp:           stateRootUpdateMessage.Timestamp,
		NearDaTransactionId: stateRootUpdateMessage.NearDaTransactionId[:],
		NearDaCommitment:    stateRootUpdateMessage.NearDaCommitment[:],
		StateRoot:           stateRootUpdateMessage.StateRoot[:],
	}

	tx := d.db.
		Where("rollup_id = ?", stateRootUpdateMessage.RollupId).
		Where("block_height = ?", stateRootUpdateMessage.BlockHeight).
		FirstOrCreate(&model)

	return &model, tx.Error
}

func (d *Database) FetchStateRootUpdate(rollupId uint32, blockHeight uint64) (*messages.StateRootUpdateMessage, error) {
	start := time.Now()
	defer func() { d.listener.OnFetch(time.Since(start)) }()

	var model models.StateRootUpdateMessage
	tx := d.db.
		Where("rollup_id = ?", rollupId).
		Where("block_height = ?", blockHeight).
		First(&model)
	if tx.Error != nil {
		return nil, tx.Error
	}

	stateRootUpdateMessage := model.ToMessage()

	return &stateRootUpdateMessage, nil
}

func (d *Database) StoreStateRootUpdateAggregation(stateRootUpdateMessage *models.StateRootUpdateMessage, aggregation messages.MessageBlsAggregation) error {
	start := time.Now()
	defer func() { d.listener.OnStore(time.Since(start)) }()

	model := models.NewMessageBlsAggregationModel(aggregation)

	err := d.db.
		Unscoped().
		Model(stateRootUpdateMessage).
		Association("Aggregation").
		Unscoped().
		Replace(&model)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) FetchStateRootUpdateAggregation(rollupId uint32, blockHeight uint64) (*messages.MessageBlsAggregation, error) {
	start := time.Now()
	defer func() { d.listener.OnFetch(time.Since(start)) }()

	var model models.StateRootUpdateMessage

	tx := d.db.
		Preload("Aggregation").
		Model(&model).
		Where("rollup_id = ?", rollupId).
		Where("block_height = ?", blockHeight).
		First(&model)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if model.Aggregation == nil {
		return nil, errors.New("aggregation not found")
	}

	aggregation := model.Aggregation.ToMessage()

	return &aggregation, nil
}

func (d *Database) StoreOperatorSetUpdate(operatorSetUpdateMessage messages.OperatorSetUpdateMessage) (*models.OperatorSetUpdateMessage, error) {
	start := time.Now()
	defer func() { d.listener.OnStore(time.Since(start)) }()

	model := models.OperatorSetUpdateMessage{
		UpdateId:  operatorSetUpdateMessage.Id,
		Timestamp: operatorSetUpdateMessage.Timestamp,
		Operators: operatorSetUpdateMessage.Operators,
	}

	tx := d.db.
		Where("update_id = ?", operatorSetUpdateMessage.Id).
		FirstOrCreate(&model)

	return &model, tx.Error
}

func (d *Database) FetchOperatorSetUpdate(id uint64) (*messages.OperatorSetUpdateMessage, error) {
	start := time.Now()
	defer func() { d.listener.OnFetch(time.Since(start)) }()

	var model models.OperatorSetUpdateMessage

	tx := d.db.
		Where("update_id = ?", id).
		First(&model)
	if tx.Error != nil {
		return nil, tx.Error
	}
	operatorSetUpdateMessage := model.ToMessage()

	return &operatorSetUpdateMessage, nil
}

func (d *Database) StoreOperatorSetUpdateAggregation(operatorSetUpdateMessage *models.OperatorSetUpdateMessage, aggregation messages.MessageBlsAggregation) error {
	start := time.Now()
	defer func() { d.listener.OnStore(time.Since(start)) }()

	model := models.NewMessageBlsAggregationModel(aggregation)

	err := d.db.
		Unscoped().
		Model(operatorSetUpdateMessage).
		Association("Aggregation").
		Unscoped().
		Replace(&model)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) FetchOperatorSetUpdateAggregation(id uint64) (*messages.MessageBlsAggregation, error) {
	start := time.Now()
	defer func() { d.listener.OnFetch(time.Since(start)) }()

	var model models.OperatorSetUpdateMessage

	tx := d.db.
		Preload("Aggregation").
		Model(&model).
		Where("update_id = ?", id).
		First(&model)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if model.Aggregation == nil {
		return nil, errors.New("aggregation not found")
	}

	aggregation := model.Aggregation.ToMessage()

	return &aggregation, nil
}

func (d *Database) FetchCheckpointMessages(fromTimestamp uint64, toTimestamp uint64) (*messages.CheckpointMessages, error) {
	if fromTimestamp > math.MaxInt64 || toTimestamp > math.MaxInt64 {
		return nil, errors.New("timestamp does not fit in int64")
	}

	if toTimestamp < fromTimestamp {
		return nil, errors.New("toTimestamp is less than fromTimestamp")
	}

	start := time.Now()
	defer func() { d.listener.OnFetch(time.Since(start)) }()

	var stateRootUpdates []models.StateRootUpdateMessage

	tx := d.db.
		Preload("Aggregation").
		Model(&models.StateRootUpdateMessage{}).
		Where("timestamp >= ?", fromTimestamp).
		Where("timestamp <= ?", toTimestamp).
		Find(&stateRootUpdates)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var operatorSetUpdates []models.OperatorSetUpdateMessage

	tx = d.db.
		Preload("Aggregation").
		Model(&models.OperatorSetUpdateMessage{}).
		Where("timestamp >= ?", fromTimestamp).
		Where("timestamp <= ?", toTimestamp).
		Find(&operatorSetUpdates)
	if tx.Error != nil {
		return nil, tx.Error
	}

	stateRootUpdateMessages := make([]messages.StateRootUpdateMessage, 0, len(stateRootUpdates))
	stateRootUpdateMessageAggregations := make([]messages.MessageBlsAggregation, 0, len(stateRootUpdates))
	operatorSetUpdateMessages := make([]messages.OperatorSetUpdateMessage, 0, len(operatorSetUpdates))
	operatorSetUpdateMessageAggregations := make([]messages.MessageBlsAggregation, 0, len(operatorSetUpdates))

	for _, stateRootUpdate := range stateRootUpdates {
		if stateRootUpdate.Aggregation == nil {
			d.db.Logger.Warn(context.Background(), "Aggregation is nil for stateRootUpdate: %v", stateRootUpdate)
			continue
		}
		agg := stateRootUpdate.Aggregation

		stateRootUpdateMessages = append(stateRootUpdateMessages, stateRootUpdate.ToMessage())
		stateRootUpdateMessageAggregations = append(stateRootUpdateMessageAggregations, agg.ToMessage())
	}

	for _, operatorSetUpdate := range operatorSetUpdates {
		if operatorSetUpdate.Aggregation == nil {
			d.db.Logger.Warn(context.Background(), "Aggregation is nil for operatorSetUpdate: %v", operatorSetUpdate)
			continue
		}
		agg := operatorSetUpdate.Aggregation

		operatorSetUpdateMessages = append(operatorSetUpdateMessages, operatorSetUpdate.ToMessage())
		operatorSetUpdateMessageAggregations = append(operatorSetUpdateMessageAggregations, agg.ToMessage())
	}

	result := &messages.CheckpointMessages{
		StateRootUpdateMessages:              stateRootUpdateMessages,
		StateRootUpdateMessageAggregations:   stateRootUpdateMessageAggregations,
		OperatorSetUpdateMessages:            operatorSetUpdateMessages,
		OperatorSetUpdateMessageAggregations: operatorSetUpdateMessageAggregations,
	}

	return result, nil
}

func (d *Database) DB() *gorm.DB {
	return d.db
}
