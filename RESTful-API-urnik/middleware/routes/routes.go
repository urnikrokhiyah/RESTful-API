package routes

import (
	"mvc/constants"
	"mvc/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("/login", controllers.LoginUserController)

	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// user with auth
	r.GET("/users/:id", controllers.GetUserDetailController)
	r.GET("/users", controllers.GetUserController)
	r.DELETE("/users/:id", controllers.DeleteUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)

	//book with auth
	r.POST("/books", controllers.CreateBookController)
	r.DELETE("/books/:id", controllers.DeleteBookController)
	r.PUT("/books/:id", controllers.UpdateBookController)

	n := e.Group("")
	// user without auth
	n.POST("/users", controllers.CreateUserController)

	// book without auth
	n.GET("/books", controllers.GetBookController)
	n.GET("/books/:id", controllers.GetOneBookController)

	// eAuth := e.Group("")
	// eAuth.Use(middleware.BasicAuth(middlewares.BasicAuthDb))

	// eAuth.DELETE("/users/:id", controllers.DeleteUserController)
	// eAuth.PUT("/users/:id", controllers.UpdateUserController)

	return e
}
