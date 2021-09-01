package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

var DB *gorm.DB

func initDB() {

	config := Config{
		DB_Username: "root",
		DB_Password: "alta123",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "alterra",
	}

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username, config.DB_Password, config.DB_Host, config.DB_Port, config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func initialMigration() {
	DB.AutoMigrate(&User{})
}

func getUsersController(c echo.Context) error {
	var users []User

	err := DB.Find(&users).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users ",
		"users":    users,
	})
}

func getOneUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user []User
	err1 := DB.Find(&user, id).Error
	if err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "user id not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get single user",
		"user":     user,
	})
}

func deleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user []User
	err1 := DB.Delete(&user, id).Error
	if err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "user id not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "deleted",
		"user id":  id,
	})
}

func updateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user []User

	tx := DB.Find(&user, id)
	if tx.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"messages": "Internal Server Error",
		})
	}
	if tx.RowsAffected > 0 {
		newUser := User{}
		c.Bind(&newUser)

		err2 := DB.Model(&user).Updates(newUser).Error
		if err2 != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"messages": "failed",
			})
		} else {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "user id updated",
				"user":     user,
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "failed to update data",
	})

}

func createUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	err := DB.Save(&user).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create new user",
		"user":     user,
	})
}

func main() {
	e := echo.New()
	e.GET("/users", getUsersController)
	e.GET("/users/:id", getOneUserController)
	e.POST("/users", createUserController)
	e.PUT("/users/:id", updateUserController)
	e.DELETE("/users/:id", deleteUserController)
	e.Start(":8000")
}

func init() {
	initDB()
	initialMigration()
}
