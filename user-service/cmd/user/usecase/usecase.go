package usecase

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"time"
	"user-service/cmd/user/services"
	"user-service/infrastructure/log"
	"user-service/models"
	"user-service/utils"
)

type UserUsecase struct {
	UserService services.UserService
	JWTSecret   string
}

func NewUserUsecase(userService services.UserService, JWTSecret string) *UserUsecase {
	return &UserUsecase{
		UserService: userService,
		JWTSecret:   JWTSecret,
	}
}

func (uc *UserUsecase) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := uc.UserService.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUsecase) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	user, err := uc.UserService.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *models.User) error {
	hPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": user.Email,
		}).Errorf("Error hashing password: %v", err)

		return err
	}

	user.Password = hPassword

	_, err = uc.UserService.CreateUser(ctx, user)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": user.Email,
		}).Errorf("Error creating user: %v", err)
		return err
	}
	return nil
}

func (uc *UserUsecase) Login(ctx context.Context, params models.LoginRequest) (string, error) {
	user, err := uc.UserService.GetUserByEmail(ctx, params.Email)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": params.Email,
		}).Errorf("Error fetching user: %v", err)
		return "", err
	}
	isMatch, err := utils.CheckPasswordHash(params.Password, user.Password)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": params.Email,
		}).Errorf("Error checking password: %v", err)
		return "", err
	}

	if !isMatch {
		return "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(uc.JWTSecret))
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": params.Email,
		}).Errorf("Error signing token: %v", err)
		return "", err
	}

	return tokenString, nil
}
