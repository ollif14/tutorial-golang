package router

import (
	"github.com/labstack/echo/v4"
	"golang-multipleDB/controller"
)

func Init() *echo.Echo{

	e := echo.New()

	e.GET("/employees", controller.GetAllEmployeesFromMultipleDB)
	e.GET("/employees/:id", controller.GetEmployeeByIdFromMultipleDB)

	return e
}
