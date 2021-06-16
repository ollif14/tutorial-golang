package main

import (
	"github.com/labstack/echo/v4"
	"golang-redis/config"
	"golang-redis/service"
	"log"
)

func main()  {
	e := echo.New()
	_, err := config.RedisClient()
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
	}

	e.POST("/student",service.Create)
	e.GET("/student", service.GetByKey)
	e.DELETE("/student", service.DeleteByKey)

	e.Logger.Fatal(e.Start(":8082"))
}
