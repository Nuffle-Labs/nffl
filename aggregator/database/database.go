package database

import (
	"errors"
	"log"
	"math"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"

	"github.com/NethermindEth/near-sffl/aggregator/database/models"
	"github.com/NethermindEth/near-sffl/core"
	"github.com/NethermindEth/near-sffl/core/types/messages"
)

type Databaser interface {
	Close() error
	StoreStateRootUpdate(stateRootUpdateMessage messages.StateRootUpdateMessage) error
	FetchStateRootUpdate(rollupId uint32, blockHeight uint64) (*messages.StateRootUpdateMessage, error)
	StoreStateRootUpdateAggregation(stateRootUpdateMessage messages.StateRootUpdateMessage, aggregation messages.MessageBlsAggregation) error
	FetchStateRootUpdateAggregation(rollupId uint32, blockHeight uint64) (*messages.MessageBlsAggregation, error)
	StoreOperatorSetUpdate(operatorSetUpdateMessage messages.OperatorSetUpdateMessage) error
	FetchOperatorSetUpdate(id uint64) (*messages.OperatorSetUpdateMessage, error)
	StoreOperatorSetUpdateAggregation(operatorSetUpdateMessage messages.OperatorSetUpdateMessage, aggregation messages.MessageBlsAggregation) error
	FetchOperatorSetUpdateAggregation(id uint64) (*messages.MessageBlsAggregation, error)
	FetchCheckpointMessages(fromTimestamp uint64, toTimestamp uint64) (*messages.CheckpointMessages, error)

	core.Metricable
}

type Database struct {
	db       *gorm.DB
	dbPath   string
	listener EventListener
}

var _ core.Metricable = (*Database)(nil)

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

	db.AutoMigrate(
		&models.MessageBlsAggregation{},
		&models.StateRootUpdateMessage{},
		&models.OperatorSetUpdateMessage{},
	)

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

func (d *Database) WithMetrics(registry *prometheus.Registry) {
	d.listener = MakeDBMetrics(registry)
}

func (d *Database) StoreStateRootUpdate(stateRootUpdateMessage messages.StateRootUpdateMessage) error {
	start := time.Now()
	defer func() { d.listener.OnStore(time.Since(start)) }()

	tx := d.db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&models.StateRootUpdateMessage{
			RollupId:            stateRootUpdateMessage.RollupId,
			BlockHeight:         stateRootUpdateMessage.BlockHeight,
			Timestamp:           stateRootUpdateMessage.Timestamp,
			NearDaTransactionId: stateRootUpdateMessage.NearDaTransactionId[:],
			NearDaCommitment:    stateRootUpdateMessage.NearDaCommitment[:],
			StateRoot:           stateRootUpdateMessage.StateRoot[:],
		})

	return tx.Error
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

func (d *Database) StoreStateRootUpdateAggregation(stateRootUpdateMessage messages.StateRootUpdateMessage, aggregation messages.MessageBlsAggregation) error {
	start := time.Now()
	defer func() { d.listener.OnStore(time.Since(start)) }()

	model := models.NewMessageBlsAggregationModel(aggregation)

	err := d.db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Model(&models.StateRootUpdateMessage{}).
		Where("rollup_id = ?", stateRootUpdateMessage.RollupId).
		Where("block_height = ?", stateRootUpdateMessage.BlockHeight).
		Association("Aggregation").
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

func (d *Database) StoreOperatorSetUpdate(operatorSetUpdateMessage messages.OperatorSetUpdateMessage) error {
	start := time.Now()
	defer func() { d.listener.OnStore(time.Since(start)) }()

	tx := d.db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&models.OperatorSetUpdateMessage{
			UpdateId:  operatorSetUpdateMessage.Id,
			Timestamp: operatorSetUpdateMessage.Timestamp,
			Operators: operatorSetUpdateMessage.Operators,
		})

	return tx.Error
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

func (d *Database) StoreOperatorSetUpdateAggregation(operatorSetUpdateMessage messages.OperatorSetUpdateMessage, aggregation messages.MessageBlsAggregation) error {
	start := time.Now()
	defer func() { d.listener.OnStore(time.Since(start)) }()

	model := models.NewMessageBlsAggregationModel(aggregation)

	err := d.db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Model(&models.OperatorSetUpdateMessage{}).
		Where("update_id = ?", operatorSetUpdateMessage.Id).
		Association("Aggregation").
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
		agg := stateRootUpdate.Aggregation

		stateRootUpdateMessages = append(stateRootUpdateMessages, stateRootUpdate.ToMessage())
		stateRootUpdateMessageAggregations = append(stateRootUpdateMessageAggregations, agg.ToMessage())
	}

	for _, operatorSetUpdate := range operatorSetUpdates {
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
