package middlewares

import (
	"mvc/config"
	"mvc/models"

	"github.com/labstack/echo/v4"
)

func BasicAuthDb(email, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.User

	err := db.Where("email = ? AND password = ?", email, password).First(&user).Error

	if err != nil {
		return false, nil
	}
	return true, nil
}
