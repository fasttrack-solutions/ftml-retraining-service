package api

import (
	"crypto/subtle"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ah *APIHandler) validateAuthenticationToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		const apiKeyHeader = "x-api-key"

		requestAPIKey := strings.ToLower(strings.TrimSpace(c.GetHeader(apiKeyHeader)))

		if len(requestAPIKey) == 0 || subtle.ConstantTimeCompare([]byte(requestAPIKey), []byte(ah.token)) != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Key"})
			return
		}

		c.Next()
	}
}
