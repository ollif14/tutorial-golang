package main

import (
	"fmt"
	"golang-multipleDB/config"
	"golang-multipleDB/router"
)

func main()  {
	fmt.Println("echo start")
	config.ConnectToMongo()
	config.ConnectToPostgre()


	e := router.Init()

	e.Logger.Fatal(e.Start(":8082"))
}
