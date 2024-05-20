package consumer

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/prometheus/client_golang/prometheus"
	rmq "github.com/rabbitmq/amqp091-go"

	"github.com/NethermindEth/near-sffl/core"
)

const (
	RECONNECT_DELAY = 3 * time.Second
	RECHANNEL_DELAY = 2 * time.Second
	EXCHANGE_NAME   = "rollup_exchange"
)

var (
	AlreadyClosedError = errors.New("Consumer connection is already closed")
)

func getQueueName(rollupId uint32, id string) string {
	return fmt.Sprintf("rollup%s-%s", strconv.FormatUint(uint64(rollupId), 10), id)
}

func getRotingKey(rollupId uint32) string {
	return fmt.Sprintf("rollup%s", strconv.FormatUint(uint64(rollupId), 10))
}

func getConsumerTag(rollupId uint32) string {
	return "operator" + strconv.FormatUint(uint64(rollupId), 10)
}

type ConsumerConfig struct {
	RollupIds []uint32
	Id        string
}

type BlockData struct {
	RollupId      uint32
	Commitment    Commitment
	TransactionId TransactionId
	Block         types.Block
}

type Consumerer interface {
	Reconnect(addr string, ctx context.Context)
	ResetChannel(conn *rmq.Connection, ctx context.Context) bool
	Close() error
	GetBlockStream() <-chan BlockData
}

type Consumer struct {
	receivedBlocksC chan BlockData
	queuesListener  *QueuesListener

	id        string
	rollupIds []uint32

	isReady           bool
	contextCancelFunc context.CancelFunc
	connection        *rmq.Connection
	connClosedErrC    <-chan *rmq.Error
	channel           *rmq.Channel
	chanClosedErrC    <-chan *rmq.Error

	logger        logging.Logger
	eventListener EventListener
}

var _ core.Metricable = (*Consumer)(nil)

func NewConsumer(config ConsumerConfig, logger logging.Logger) *Consumer {
	consumer := Consumer{
		id:              config.Id,
		rollupIds:       config.RollupIds,
		receivedBlocksC: make(chan BlockData),
		logger:          logger,
		eventListener:   &SelectiveListener{},
	}

	return &consumer
}

func (consumer *Consumer) EnableMetrics(registry *prometheus.Registry) error {
	eventListener, err := MakeConsumerMetrics(registry)
	if err != nil {
		return err
	}

	consumer.eventListener = eventListener
	return nil
}

func (consumer *Consumer) Start(ctx context.Context, addr string) {
	ctx, cancel := context.WithCancel(ctx)
	consumer.contextCancelFunc = cancel

	go consumer.reconnect(ctx, addr)
}

func (consumer *Consumer) reconnect(ctx context.Context, addr string) {
	for {
		consumer.logger.Info("Reconnecting...")

		consumer.isReady = false
		conn, err := consumer.connect(addr)
		if err != nil {
			consumer.logger.Warn("Connection setup failed", "err", err)

			select {
			case <-ctx.Done():
				consumer.logger.Info("Consumer context canceled")
				return
			case <-time.After(RECONNECT_DELAY):
			}

			continue
		}

		if done := consumer.ResetChannel(ctx, conn); done {
			return
		}

		consumer.logger.Info("Connected")

		select {
		case <-ctx.Done():
			consumer.logger.Info("Consumer context canceled")
			// deref cancel smth?
			return

		case err := <-consumer.connClosedErrC:
			if !err.Recover {
				consumer.logger.Error("Can't recover connection", "err", err)
				break
			}

			consumer.logger.Warn("Recovering connection, closed with:", "err", err)

		case err := <-consumer.chanClosedErrC:
			if !err.Recover {
				consumer.logger.Error("Can't recover connection", "err", err)
				break
			}

			consumer.logger.Warn("Reconnecting channel, closed with:", "err", err)
		}
	}
}

func (consumer *Consumer) connect(addr string) (*rmq.Connection, error) {
	conn, err := rmq.Dial(addr)
	if err != nil {
		return nil, err
	}

	consumer.changeConnection(conn)
	return conn, nil
}

func (consumer *Consumer) changeConnection(conn *rmq.Connection) {
	consumer.connection = conn

	connClosedErrC := make(chan *rmq.Error)
	consumer.connClosedErrC = conn.NotifyClose(connClosedErrC)
}

func (consumer *Consumer) ResetChannel(ctx context.Context, conn *rmq.Connection) bool {
	for {
		consumer.isReady = false

		err := consumer.setupChannel(ctx, conn)
		if err != nil {
			consumer.logger.Warn("Channel setup failed", "err", err)

			select {
			case <-ctx.Done():
				consumer.logger.Info("Consumer context canceled")
				return true

			case rmqError := <-consumer.connClosedErrC:
				if rmqError.Recover {
					consumer.logger.Error("Can't recover connection", "err", err)
					return true
				}

				consumer.logger.Warn("Recovering connection, closed with:", "err", err)
				return false
			case <-time.After(RECHANNEL_DELAY):
			}

			continue
		}

		return false
	}
}

func (consumer *Consumer) setupChannel(ctx context.Context, conn *rmq.Connection) error {
	channel, err := conn.Channel()
	if err != nil {
		return err
	}

	listener := NewQueuesListener(consumer.receivedBlocksC, consumer.eventListener, consumer.logger)
	for _, rollupId := range consumer.rollupIds {
		queue, err := channel.QueueDeclare(getQueueName(rollupId, consumer.id), true, false, false, false, nil)
		if err != nil {
			return err
		}

		err = channel.QueueBind(queue.Name, getRotingKey(rollupId), EXCHANGE_NAME, false, nil)
		if err != nil {
			return err
		}

		consumerTag := getConsumerTag(rollupId)
		rollupDataC, err := channel.Consume(
			queue.Name,
			consumerTag,
			false,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			return err
		}

		err = listener.Add(ctx, rollupId, rollupDataC)
		if err != nil {
			return err
		}
	}

	consumer.queuesListener = listener
	consumer.changeChannel(channel)
	consumer.isReady = true
	return nil
}

func (consumer *Consumer) changeChannel(channel *rmq.Channel) {
	consumer.channel = channel

	chanClosedErrC := make(chan *rmq.Error)
	consumer.chanClosedErrC = channel.NotifyClose(chanClosedErrC)
}

func (consumer *Consumer) Close() error {
	if !consumer.isReady {
		return AlreadyClosedError
	}

	// shut down goroutines
	if consumer.contextCancelFunc != nil {
		consumer.contextCancelFunc()
	}

	err := consumer.channel.Close()
	if err != nil {
		return err
	}

	err = consumer.connection.Close()
	if err != nil {
		return err
	}

	consumer.isReady = false
	return nil
}

func (consumer *Consumer) GetBlockStream() <-chan BlockData {
	return consumer.receivedBlocksC
}
