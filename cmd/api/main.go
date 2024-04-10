package main

import (
	"github.com/fasttrack-solutions/ftml-retraining-service/internal/app/api"
	"github.com/fasttrack-solutions/ftml-retraining-service/internal/config"
	"github.com/fasttrack-solutions/ftml-retraining-service/internal/logger"
)

func main() {
	logger := logger.NewLogger()
	logger.Info("Starting FTML Retraining API")

	apiHandler := api.NewAPI(&logger, *config.APIKey)
	apiHandler.InitializeRoutes()
}
