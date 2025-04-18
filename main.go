package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rubenbuelvas/garage-be/src/api/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	engine := gin.Default()
	routes.SetupRoutes(engine)
	engine.Run(port)
}
