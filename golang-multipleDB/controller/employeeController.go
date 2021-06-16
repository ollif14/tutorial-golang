package controller

import (
	"github.com/labstack/echo/v4"
	"golang-multipleDB/services"
	"net/http"
)

func GetAllEmployeesFromMultipleDB(c echo.Context) error {
	c.JSON(http.StatusCreated, "Mongo DB : ")
	services.GetAllEmployeesfromMongo(c)
	c.JSON(http.StatusCreated, "Postgre DB : ")
	services.GetAllEmployeesfromPg(c)

	return nil
}

func GetEmployeeByIdFromMultipleDB(c echo.Context) error {
	c.JSON(http.StatusCreated, "Mongo DB : ")
	services.GetEmployeeByIdMongo(c)
	c.JSON(http.StatusCreated, "Postgre DB : ")
	services.GetEmployeeByIdPostgre(c)

	return nil
}
