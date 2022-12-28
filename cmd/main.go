package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"place4live/internal/infrastructure/numbeo"
)

func getCity(c *gin.Context) {
	n := c.Param("name")

	log.Printf("City name: %s\n", n)

	city := <-numbeo.GetCity(n)

	c.JSON(http.StatusOK, city)
}

func main() {
	r := gin.New()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/city/:name", getCity)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
