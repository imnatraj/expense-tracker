package service

import (
	"errors"
	"fmt"
	"imnatraj/expense-tracker/models"

	"gorm.io/gorm"
)

// db related actions will be performed here

type user struct {
	*models.HandlerCtx
}

func User(ctx *models.HandlerCtx) *user {
	return &user{ctx}
}

func (ctx *user) Create(payload *models.User) error {
	// Check if an existing user with the same username and email ID exists
	var existingUser models.User

	if err := ctx.DB.Where("username = ?", payload.Username).First(&existingUser).Error; err == nil {
		return errors.New("Username already exists")
	}

	// Reset existingUser before the second query
	existingUser = models.User{}
	if err := ctx.DB.Where("email = ?", payload.Email).First(&existingUser).Error; err == nil {
		return errors.New("User with this email already exists")
	}

	return ctx.DB.Create(payload).Error
}

func (ctx *user) GetOne(id uint) (user *models.User, err error) {
	user = &models.User{}
	if err := ctx.DB.Where("id = ?", id).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("User with Id %d not found", id)
		}
		return nil, err
	}
	return user, nil
}

func (ctx *user) Update(id uint, payload *models.User) error {
	var existingUser models.User
	err := ctx.DB.First(&existingUser, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If the error is "record not found", return a custom error
			return fmt.Errorf("User with Id %d not found", id)
		}
		// If the error is not "record not found", return the original error
		return err
	}
	return ctx.DB.Model(&existingUser).Updates(payload).Error
}

func (ctx *user) GetMany() (users []models.User, err error) {
	err = ctx.DB.Find(&users).Error
	return
}
