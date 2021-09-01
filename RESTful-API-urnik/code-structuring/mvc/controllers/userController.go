package controllers

import (
	"mvc/lib/database"
	"mvc/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	users, err := database.GetUser()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    users,
	})
}

func GetSingleController(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	user, _, err := database.GetSingleUser(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    user,
	})
}

func CreateUserController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	users, err := database.CreateUser(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    users,
	})
}

func DeleteUserController(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	_, err := database.DeleteUser(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "user id not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "deleted",
		"user id":  id,
	})
}

func UpdateUserController(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	user, row, err := database.GetSingleUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"messages": "Internal Server Error",
		})
	}

	if row > 0 {
		newUser := models.User{}
		c.Bind(&newUser)

		_, err := database.UpdateUser(user, newUser)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"messages": "failed",
			})
		} else {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "user is updated",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "failed to update data",
	})
}
