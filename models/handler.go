package models

import (
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type HandlerCtx struct {
	DB            *gorm.DB
	AnotherDB     *gorm.DB
	UserSessionDB *redis.Client
	OtpSessionDB  *redis.Client
	
}
