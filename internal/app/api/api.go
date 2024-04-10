package api

import (
	"github.com/fasttrack-solutions/ftml-retraining-service/internal/logger"
	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	token  string
	router *gin.Engine
}

func NewAPI(logger *logger.Logger, token string) APIHandler {
	r := gin.New()

	r.Use(logger.StartGinLogger())
	r.Use(gin.Recovery())

	return APIHandler{
		token:  token,
		router: r,
	}
}
