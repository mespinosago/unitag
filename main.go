package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mespinosago/unitag/internal/handler"
	"github.com/mespinosago/unitag/internal/service"
)

func main() {
	h := handler.NewHandler(service.NewService())

	router := gin.Default()
	// Register routes
	router.GET("/:code", h.GetURL)
	// Start the server
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
