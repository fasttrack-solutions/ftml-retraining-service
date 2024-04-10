package queue

import (
	"fmt"
	"log"

	msdk "github.com/fasttrack-solutions/message-sdk"
	"github.com/fasttrack-solutions/message-sdk/brokers"
	"github.com/fasttrack-solutions/message-sdk/brokers/rabbitmq"
	topic "github.com/fasttrack-solutions/queue-topic-management"
)

type RabbitService struct {
	brandID  int
	consumer brokers.Consumer
}

// NewRabbitService returns new NewRabbitService.
func NewRabbitService(
	brandName string,
	consumerTags string,
	queue string,
	serviceName string,
	uri string,
	brandID int,
	prefetch int,
	events []topic.Bindable,
) (*RabbitService, error) {
	brokerSettings := rabbitmq.BrokerSettings{
		BrandName:            brandName,
		ServiceName:          serviceName,
		ConnectionURL:        uri,
		PrefetchSize:         prefetch,
		Queue:                queue,
		ConsumerTags:         consumerTags,
		Events:               events,
		AckArgs:              nil,
		ShouldCreateBindings: true,
	}

	consumer, err := msdk.NewConsumer(rabbitmq.Name, brokerSettings)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize broker consumer: %v", err)
	}

	return &RabbitService{
		brandID:  brandID,
		consumer: consumer,
	}, nil
}

func (r *RabbitService) StartAsync() {
	consumerSettings := rabbitmq.ConsumerSettings{
		Idempotency: nil,
		CloseHandler: func(e error) {
			log.Fatalf("amqp connection closed: %v", e)
		},
		OnDisconnect: func() {
			log.Fatal("amqp connection has been interrupted")
		},
	}

	r.consumer.StartAsync(consumerSettings, r.handler)
}

func (r *RabbitService) handler(msg []byte, headers brokers.Headers, properties brokers.Properties, ack func() error) error {
	return ack()
}
