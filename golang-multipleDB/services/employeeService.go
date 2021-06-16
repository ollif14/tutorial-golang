package services

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang-multipleDB/config"
	"golang-multipleDB/model"
	"net/http"
	"strconv"
)

func GetAllEmployeesfromPg(c echo.Context) error  {
	var e []model.Employees
	config.PgDB.Find(&e)
	return c.JSON(http.StatusCreated, e)
}

func GetAllEmployeesfromMongo(c echo.Context) error{
	result, err := config.MongoCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.JSON(http.StatusCreated, err)
	}

	var e[]bson.M
	if err = result.All(context.Background(), &e); err != nil {
		return c.JSON(http.StatusCreated, err)
	}
	return c.JSON(http.StatusCreated, e)
}

func GetEmployeeByIdPostgre(c echo.Context) error {
	var e model.Employees
	if err := config.PgDB.Where("id = ?", c.Param("id")).First(&e).Error; err != nil {
		return c.JSON(http.StatusCreated, "could not find data in postgre")
	}
	return c.JSON(http.StatusCreated, e)
}

func GetEmployeeByIdMongo(c echo.Context) error {
	var e bson.M
	id := c.Param("id")
	idInt, _ := strconv.ParseInt(id, 10, 64)
	result := config.MongoCollection.FindOne(context.Background(), bson.M{"e_id" : idInt})
	if err := result.Decode(&e); err != nil {
		return c.JSON(http.StatusCreated, "could not find data in mongodb")
	}
	return c.JSON(http.StatusCreated, e)
}
