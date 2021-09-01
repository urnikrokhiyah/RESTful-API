package middlewares

import "github.com/labstack/echo/v4"

func BasicAuth(email string, password string, c echo.Context) (bool, error) {
	if email == "admin" && password == "admin" {
		return true, nil
	}
	return false, nil
}
