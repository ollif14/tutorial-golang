package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang-kafka-2/model"
	producer2 "golang-kafka-2/producer"
	"net/http"
)

func main()  {
	e := echo.New()

	producer, err := producer2.NewProducer()
	if err != nil {
		fmt.Println("Could not create producer: ", err)
	}

	e.POST("/employees", func(context echo.Context) error {
		var req model.Employees
		if err := context.Bind(&req); err != nil{
			return context.JSON(http.StatusCreated, err)
		}
		msg := producer2.Produce("test-topic-1", req)
		_, _, err := producer.SendMessage(msg)
		if err != nil {
			context.JSON(http.StatusCreated, err)
		}
		return context.JSON(http.StatusCreated, msg)
	})
	e.Logger.Fatal(e.Start(":9093"))
}

