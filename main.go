package main

import (
	"catering-api/helpers"
	"catering-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	db := helpers.DatabaseConnect()
	myApp := echo.New()

	routes.EmployeeRoutes(myApp, db)

	myApp.Logger.Fatal(myApp.Start(":8005"))
}
