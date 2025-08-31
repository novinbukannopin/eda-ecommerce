package routes

import (
	"github.com/gin-gonic/gin"
	"product-service/middleware"
)

func SetupRoutes(router *gin.Engine, h handler.UserHandler, jwtSecret string) {
	router.Use(middleware.RequestLog())
}
