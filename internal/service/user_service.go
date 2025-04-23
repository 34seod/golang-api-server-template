package service

import (
	"errors"

	"golang-api-server-template/internal/dto"
	"golang-api-server-template/internal/model"
	"golang-api-server-template/internal/repository"
	"golang-api-server-template/tools"

	"gorm.io/gorm"
)

func UserFindByID(userDto *dto.UserIdFromUri) (*model.User, error) {
	var user *model.User

	user, err := repository.UserFindByID(userDto)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tools.PrintTrace()
		return user, err
	}

	return user, nil
}

func UserFindAll() ([]model.User, error) {
	var users []model.User

	users, err := repository.UserFindAll()
	if err != nil {
		tools.PrintTrace()
		return users, err
	}

	return users, nil
}

func UserCreate(user *model.User) error {
	if err := repository.UserCreate(user); err != nil {
		tools.PrintTrace()
		return err
	}

	return nil
}

func UserUpdate(userDto *dto.UserBodyFromUpdateRequest) error {
	if err := repository.UserUpdate(userDto); err != nil {
		tools.PrintTrace()
		return err
	}

	return nil
}

func UserDelete(userDto *dto.UserIdFromUri) error {
	if err := repository.UserDelete(userDto); err != nil {
		tools.PrintTrace()
		return err
	}

	return nil
}
