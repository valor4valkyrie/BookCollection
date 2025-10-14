package main

import (
	"bookcollection.com/rest-api/db"
	"bookcollection.com/rest-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	db.InitDB()
	e := echo.New()
	e.Server.Addr = ":8080"
	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
