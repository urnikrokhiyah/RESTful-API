package database

import (
	"mvc/config"
	"mvc/models"
)

func GetUser() (interface{}, error) {
	var users []models.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetSingleUser(id int) (interface{}, int, error) {
	var user models.User

	tx := config.DB.Find(&user, id)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	return user, 1, nil
}

func CreateUser(user *models.User) (interface{}, error) {

	err := config.DB.Create(&user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	var user models.User

	err := config.DB.Delete(&user, id).Error

	if err != nil {
		return nil, err
	}

	return "deleted", nil
}

func UpdateUser(user interface{}, newUser models.User) (interface{}, error) {
	tx := config.DB.Model(user).Updates(newUser)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return "updated", nil
}
