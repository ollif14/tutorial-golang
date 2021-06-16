package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang-mongodb/controller"
)

func main() {
	fmt.Println("GOLANG_MONGODB")

	e := echo.New()

	e.GET("/employees", controller.GetAllEmployees)
	e.GET("/employees/:id", controller.GetEmployeeById)
	e.POST("/employees", controller.CreateEmployee)
	e.PUT("/employees/:id", controller.UpdateEmployee)
	e.DELETE("/employees/:id", controller.DeleteEmployee)

	e.Logger.Fatal(e.Start(":9092"))
}
