package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
	Redis    *redis.Client
}

func NewUserRepository(db *gorm.DB, redisClient *redis.Client) *UserRepository {
	return &UserRepository{
		Database: db,
		Redis:    redisClient,
	}
}
