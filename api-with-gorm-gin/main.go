package main

import (
	"api-with-gorm/config"
	"api-with-gorm/controller"
	"api-with-gorm/docs"
	_ "api-with-gorm/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	fmt.Println("TEST MAIN")
	r := gin.Default()

	config.ConnectToDB()

	r.GET("/employees", controller.GetAllEmployees)
	r.GET("/employees/:id", controller.GetEmployeeById)
	r.POST("/employees", controller.CreateEmployee)
	r.PUT("/employees/:id", controller.UpdateEmployee)
	r.DELETE("/employees/:id", controller.DeleteEmployee)

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
