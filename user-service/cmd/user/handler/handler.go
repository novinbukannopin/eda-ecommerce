package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"user-service/cmd/user/usecase"
	"user-service/infrastructure/log"
	"user-service/models"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	var params models.LoginRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Logger.Info("Error Invalid Request: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(params.Password) < 8 {
		log.Logger.Info("Password is too short")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password must be at least 8 characters",
		})
		return
	}

	token, err := h.UserUsecase.Login(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	userIDStr, isExist := c.Get("user_id")
	if !isExist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User ID not found in token",
		})
		return
	}

	userID, ok := userIDStr.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid user ID in token",
		})
		return
	}

	user, err := h.UserUsecase.GetUserByID(c.Request.Context(), int64(userID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}
		log.Logger.Error("Error fetching user info: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})

}

func (h *UserHandler) Register(c *gin.Context) {
	var params models.RegisterRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Logger.Info("Error Invalid Request: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(params.Password) < 8 {
		log.Logger.Info("Password is too short")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password must be at least 8 characters",
		})
		return
	}

	if params.Password != params.ConfirmPassword {
		log.Logger.Info("Password confirmation does not match")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password confirmation does not match",
		})
		return
	}

	user, err := h.UserUsecase.GetUserByEmail(c, params.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) { // or your custom "not found" error
			log.Logger.Error("Error checking existing user: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}
	} else if user.ID != 0 {
		log.Logger.Info("Email already exists: ", params.Email)
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email already registered",
		})
		return
	}

	err = h.UserUsecase.CreateUser(c, &models.User{
		Name:     params.Name,
		Email:    params.Email,
		Password: params.Password,
	})

	if err != nil {
		log.Logger.Info("Error creating user: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

func (h *UserHandler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
