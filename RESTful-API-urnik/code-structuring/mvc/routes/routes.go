package routes

import (
	"mvc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetUserController)
	e.GET("/users/:id", controllers.GetSingleController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.POST("/users", controllers.CreateUserController)

	e.GET("/books", controllers.GetBookController)
	e.GET("/books/:id", controllers.GetOneBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	return e
}
