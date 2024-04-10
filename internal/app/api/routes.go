package api

import (
	"net/http"

	"github.com/fasttrack-solutions/ftml-retraining-service/internal/app/retrain"
	"github.com/gin-gonic/gin"
)

func (ah *APIHandler) InitializeRoutes() {
	ah.router.Use(ah.validateAuthenticationToken())

	setupHealthCheck(ah.router, "/health")
	ah.router.GET("/retrain", ah.StartManualRetraining)
}

func (ah *APIHandler) StartManualRetraining(c *gin.Context) {
	err := retrain.StartManualRetraining()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func setupHealthCheck(r *gin.Engine, e string) {
	r.GET(e, func(c *gin.Context) {
		c.Status(200)
		c.Abort()
	})
}
