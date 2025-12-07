package main

import (
	"os"

	"bookcollection.com/rest-api/db"
	"bookcollection.com/rest-api/gui"
	"bookcollection.com/rest-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	db.InitDB()
	e := echo.New()
	e.Server.Addr = ":8080"
	routes.RegisterRoutes(e)
	gui.StartGui()
	e.Logger.Info(e.Start(":8080"))
	os.Exit(0)
}
