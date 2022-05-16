package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricHandler(r *gin.Engine) {
	r.GET("/metrics", func(c *gin.Context) {
		handler := promhttp.Handler()
		handler.ServeHTTP(c.Writer, c.Request)
	})
}

func HealthHandlers(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{"status": http.StatusOK, "version": "v0.0.1"})
	})
}
