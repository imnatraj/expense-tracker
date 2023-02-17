package models

import (
	"gorm.io/gorm"
)

type HandlerCtx struct {
	DB *gorm.DB
}
