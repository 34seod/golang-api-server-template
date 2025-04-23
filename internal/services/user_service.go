package services

import (
	"errors"

	"golang-api-server-template/internal/dto"
	"golang-api-server-template/internal/models"
	"golang-api-server-template/internal/repositories"
	"golang-api-server-template/tools"

	"gorm.io/gorm"
)

func UserFindByID(userDto *dto.UserIdFromUri) (*models.User, error) {
	var user *models.User

	user, err := repositories.UserFindByID(userDto)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tools.PrintTrace()
		return user, err
	}

	return user, nil
}

func UserFindAll() ([]models.User, error) {
	var users []models.User

	users, err := repositories.UserFindAll()
	if err != nil {
		tools.PrintTrace()
		return users, err
	}

	return users, nil
}

func UserCreate(user *models.User) error {
	if err := repositories.UserCreate(user); err != nil {
		tools.PrintTrace()
		return err
	}

	return nil
}

func UserUpdate(userDto *dto.UserBodyFromUpdateRequest) error {
	if err := repositories.UserUpdate(userDto); err != nil {
		tools.PrintTrace()
		return err
	}

	return nil
}

func UserDelete(userDto *dto.UserIdFromUri) error {
	if err := repositories.UserDelete(userDto); err != nil {
		tools.PrintTrace()
		return err
	}

	return nil
}
