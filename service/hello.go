package service

import (
	"github.com/natraj/expense-tracker/models"
)

// db related actions will be performed here

type hello struct {
	*models.HandlerCtx
}

func Hello(ctx *models.HandlerCtx) *hello {
	return &hello{ctx}
}

func (ctx *hello) Create(payload *models.Hello) error {
	return ctx.DB.Create(payload).Error
}

func (ctx *hello) GetOne(id uint) (hello *models.Hello, err error) {
	err = ctx.DB.Table("hellos").Where(`id = ?`, id).Scan(&hello).Error
	return
}

func (ctx *hello) Update(id uint, payload *models.Hello) error {
	return ctx.DB.Model(&models.Hello{}).Where("id = ?", id).Updates(payload).Error
}

func (ctx *hello) GetMany() (hellos []models.Hello, err error) {
	err = ctx.DB.Find(&hellos).Error
	return
}
