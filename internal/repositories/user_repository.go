package repositories

import (
	. "golang-api-server-template/configs"
	"golang-api-server-template/internal/dto"
	model "golang-api-server-template/internal/models"
	"golang-api-server-template/tools"
)

func UserFindByID(userDto *dto.UserIdFromUri) (*model.User, error) {
	var user model.User

	if err := DB.Where("id = ?", userDto.ID).Take(&user).Error; err != nil {
		tools.PrintTrace()
		return nil, err
	}

	return &user, nil
}

func UserFindAll() ([]model.User, error) {
	var users []model.User

	if err := DB.Find(&users).Error; err != nil {
		tools.PrintTrace()
		return nil, err
	}

	return users, nil
}

func UserCreate(user *model.User) error {
	if err := DB.Create(user).Error; err != nil {
		tools.PrintTrace()
		return err
	}
	return nil
}

func UserUpdate(userDto *dto.UserBodyFromUpdateRequest) error {
	var user model.User

	if err := DB.Where("id = ?", userDto.ID).Take(&user).Error; err != nil {
		tools.PrintTrace()
		return err
	}

	if err := DB.Model(&user).Where("id = ?", userDto.ID).Updates(userDto).Error; err != nil {
		tools.PrintTrace()
		return err
	}

	return nil
}

func UserDelete(userDto *dto.UserIdFromUri) error {
	var user model.User

	if err := DB.Where("id = ?", userDto.ID).Take(&user).Error; err != nil {
		tools.PrintTrace()
		return err
	}

	if err := DB.Delete(&user).Error; err != nil {
		tools.PrintTrace()
		return err
	}

	return nil
}
