package main

import (
	"fmt"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang-echo/dbconfig"
	_ "golang-echo/docs"
	"golang-echo/route"
)

func main() {
	fmt.Println("echo start")
	dbconfig.ConnectToDB()
	e := route.Init()

	//for swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))

}
