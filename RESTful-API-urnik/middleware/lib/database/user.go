package database

import (
	"mvc/config"
	"mvc/middlewares"
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

func UpdateUser(id int, newUser *models.User) (interface{}, int, error) {
	var user models.User
	tx1 := config.DB.Find(&user, id)

	if tx1.Error != nil {
		return nil, 0, tx1.Error
	}

	if tx1.RowsAffected > 0 {
		tx := config.DB.Model(&user).Updates(newUser)

		if tx.Error != nil {
			return nil, 0, tx.Error
		}

		return user, 1, nil
	}
	return nil, 0, nil
}

func LoginUser(user *models.User) (interface{}, error) {
	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error

	if err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	err1 := config.DB.Save(user).Error

	if err1 != nil {
		return nil, err
	}
	return user, nil
}

func GetUserDetail(userId int) (interface{}, error) {
	var user models.User
	err := config.DB.Find(&user, userId).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}
