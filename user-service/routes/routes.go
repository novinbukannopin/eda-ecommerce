package routes

import (
	"github.com/gin-gonic/gin"
	"user-service/cmd/user/handler"
	"user-service/middleware"
)

func SetupRoutes(router *gin.Engine, h handler.UserHandler, jwtSecret string) {
	router.Use(middleware.RequestLog())
	router.GET("/ping", h.Ping)
	router.POST("/v1/register", h.Register)
	router.POST("/v1/login", h.Login)

	authMiddleware := middleware.AuthMiddleware(jwtSecret)
	private := router.Group("/api").Use(authMiddleware)
	private.GET("/v1/user_info", h.GetUserInfo)
}
