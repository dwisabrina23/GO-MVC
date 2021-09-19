package database

import (
	"book-rentals/config"
	"book-rentals/model"
)

func GetUsers() (*[]model.User, error) {
	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return &[]model.User{}, err
	}
	return &users, nil
}

func GetByIDUser(id int) (*model.User, error) {
	var user model.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return &model.User{}, err
	}
	return &user, nil
}

func AddNewUser(user model.User) (*model.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return &model.User{}, err
	}
	return &user, nil
}

func UpdateUser(id int, userData model.User) (*model.User, error) {
	var user model.User

	if err := config.DB.First(&user, id).Updates(&userData).Error; err != nil {
		return &model.User{}, err
	}
	return &user, nil
}

func DeleteUser(id int, user model.User) (string, error) {
	if err := config.DB.First(&user, id).Delete(&user).Error; err != nil {
		return "", err
	}
	return "deleted", nil
}
