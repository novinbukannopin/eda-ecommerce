package services

import (
	"context"
	"user-service/cmd/user/repository"
	"user-service/models"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

func (svc *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := svc.UserRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (svc *UserService) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	user, err := svc.UserRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (svc *UserService) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	id, err := svc.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
