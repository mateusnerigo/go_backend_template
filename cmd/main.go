package main

import (
	"backend/internal/delivery/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.RegisterRoutes(server)

	log.Fatal(server.Run(":8000"))
}
