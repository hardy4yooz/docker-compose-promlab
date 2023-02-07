package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "requests_total",
		Help: "Number of requests processed, partitioned by status code and HTTP method.",
	},
	[]string{"code", "method"},
)

func init() {
	prometheus.MustRegister(requestCounter)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		// Increase request count
		requestCounter.With(prometheus.Labels{"code": "200", "method": "GET"}).Inc()
		c.JSON(http.StatusOK, gin.H{
			"Status": "passing",
			"Output": "Everything is okay",
		})
	})

	r.GET("/metrics", func(c *gin.Context) {
		h := promhttp.Handler()
		h.ServeHTTP(c.Writer, c.Request)
	})

	r.Run(":5001")

}
