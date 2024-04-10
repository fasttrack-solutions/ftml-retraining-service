package main

import (
	"log"

	"github.com/fasttrack-solutions/ftml-retraining-service/internal/config"
	"github.com/fasttrack-solutions/ftml-retraining-service/internal/database"
	"github.com/fasttrack-solutions/ftml-retraining-service/internal/logger"
	"github.com/fasttrack-solutions/ftml-retraining-service/internal/queue"
	topic "github.com/fasttrack-solutions/queue-topic-management"
)

func main() {
	logger := logger.NewLogger()
	logger.Info("Starting FTML Retraining Service")

	events := []topic.Bindable{
		topic.Time,
	}

	db := database.NewDataManager()
	db.Connect(*config.DatabaseConnStr)
	defer db.Close()

	msgConsumer, err := queue.NewRabbitService(
		*config.Brand,
		*config.ConsumerTags,
		*config.AMQPQueue,
		*config.ServiceName,
		*config.AMQPConnectionStr,
		*config.BrandID,
		*config.AMQPPrefetchSize,
		events,
	)
	if err != nil {
		log.Fatalf("Failed to start broker service: %v", err)
	}

	msgConsumer.StartAsync()
}
