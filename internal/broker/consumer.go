package broker

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-runner-api/internal/models"
)

type Consumer interface {
	Init(brokers []string) error
	Subscribe(topic string) error
	Close()
}

// NewConsumer constructor
func NewConsumer() Consumer {
	return &consumer{}
}

type consumer struct {
	sc sarama.Consumer
}

// Init initializes consumer
func (c *consumer) Init(brokers []string) error {
	cons, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		return err
	}
	c.sc = cons
	return nil
}

// Subscribe subscribes to topic
func (c *consumer) Subscribe(topic string) error {
	partitionList, err := c.sc.Partitions(topic)
	if err != nil {
		return err
	}

	initialOffset := sarama.OffsetOldest
	for _, partition := range partitionList {
		pc, err := c.sc.ConsumePartition(topic, partition, initialOffset)
		if err != nil {
			return err
		}

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				c.handleMessage(message)
			}
		}(pc)
	}

	return nil
}

// Close closes connection
func (c *consumer) Close() {
	c.sc.Close()
}

// handleMessage processes message from topic
func (c *consumer) handleMessage(message *sarama.ConsumerMessage) {
	log.Info().Msgf("received message: %v", string(message.Value))

	var runnerEvent models.RunnerEvent
	if err := json.Unmarshal(message.Value, &runnerEvent); err != nil {
		log.Error().AnErr("failed to parse message", err).Send()
		return
	}

	log.Info().Str("event", runnerEvent.String()).Fields(runnerEvent.Body).Send()
}
