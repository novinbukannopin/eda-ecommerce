package main

import (
	"github.com/gin-gonic/gin"
	"user-service/cmd/user/handler"
	"user-service/cmd/user/repository"
	"user-service/cmd/user/resources"
	"user-service/cmd/user/services"
	"user-service/cmd/user/usecase"
	"user-service/config"
	"user-service/infrastructure/log"
	"user-service/routes"
)

func main() {
	cfg := config.LoadConfig()
	redis := resources.InitRedis(&cfg)
	db := resources.InitDatabase(&cfg)

	log.SetupLogger()

	userRepository := repository.NewUserRepository(db, redis)
	userServices := services.NewUserService(*userRepository)
	userUsecase := usecase.NewUserUsecase(*userServices, cfg.Secret.JWTSecret)
	userHandler := handler.NewUserHandler(*userUsecase)

	port := cfg.App.Port

	router := gin.Default()
	routes.SetupRoutes(router, *userHandler, cfg.Secret.JWTSecret)

	err := router.Run(":" + port)
	if err != nil {
		return
	}

	log.Logger.Printf("Server running on port %s", port)

}
