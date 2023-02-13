package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// router for http sd
	r.GET("/targets", func(c *gin.Context) {
		targets := []struct {
			Targets []string          `json:"targets"`
			Labels  map[string]string `json:"labels"`
		}{
			{
				Targets: []string{"http_service2:5002"},
				Labels:  map[string]string{"job": "go-gin-demo2"},
			},
		}

		c.JSON(http.StatusOK, targets)
	})

	r.Run(":5000")
}
