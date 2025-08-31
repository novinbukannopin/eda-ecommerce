package main

import (
	"product-service/cmd/product/resources"
	"product-service/config"
	"product-service/infrastructure/log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	resources.InitRedis(&cfg)
	resources.InitDatabase(&cfg)

	log.SetupLogger()

	port := cfg.App.Port

	router := gin.Default()
	router.Run(":" + port)
	log.Logger.Printf("Server running on port %s", port)
}
