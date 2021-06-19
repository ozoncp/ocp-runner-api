package broker

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-runner-api/internal/models"
)

type Producer interface {
	Init(brokers []string) error
	SendMessage(topic string, event *models.RunnerEvent) error
	Close()
}

// NewProducer constructor
func NewProducer() Producer {
	return &producer{}
}

type producer struct {
	sp sarama.AsyncProducer
}

// Init initialize broker
func (p *producer) Init(brokers []string) error {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	prod, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return err
	}
	p.sp = prod

	return nil
}

// SendMessage processes message sending
func (p *producer) SendMessage(topic string, event *models.RunnerEvent) error {
	message, err := p.prepareMessage(topic, event)
	if err != nil {
		return err
	}

	p.sp.Input() <- message

	select {
	case msg := <-p.sp.Successes():
		log.Info().Msgf("message successfully published: %v", msg.Value)
		return nil
	case err := <-p.sp.Errors():
		log.Error().Err(err).Send()
		return err
	}
}

// prepareMessage prepares runner event message to sending
func (p *producer) prepareMessage(topic string, event *models.RunnerEvent) (*sarama.ProducerMessage, error) {
	data, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	result := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(data),
	}

	return result, nil
}

// Close closes sarama broker
func (p *producer) Close() {
	p.sp.Close()
}
