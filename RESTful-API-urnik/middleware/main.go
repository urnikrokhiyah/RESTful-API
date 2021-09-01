package main

import (
	"mvc/config"
	m "mvc/middlewares"
	"mvc/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
