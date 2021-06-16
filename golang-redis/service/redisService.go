package service

import (
	"github.com/labstack/echo/v4"
	"golang-redis/model"
	"golang-redis/repository"
	"net/http"
)

func Create(c echo.Context) error {
	var student model.Student
	if err := c.Bind(&student); err != nil{
		c.JSON(http.StatusCreated, err)
	}

	if err := repository.Save(&student); err != nil {
		c.JSON(http.StatusCreated, err)
	}
	return c.JSON(http.StatusCreated, student)
}

func GetByKey(c echo.Context)error{
	var key = "student"
	result, err := repository.GetBy(key)
	if err != nil {
		c.JSON(http.StatusCreated, err)
	}
	return c.JSON(http.StatusCreated, result)
}

func DeleteByKey(c echo.Context) error{
	var key = "student"
	err := repository.Delete(key)
	if err != nil {
		c.JSON(http.StatusCreated, err)
	}
	return c.JSON(http.StatusCreated, true)
}