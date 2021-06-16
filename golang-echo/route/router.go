package route

import (
	"github.com/labstack/echo/v4"
	"golang-echo/handler"
	"golang-echo/service"
	"net/http"
)

func Init() *echo.Echo{

	e := echo.New()

	var jwtService service.JWTService = service.JWTAuthService()
	var tokenHandler handler.TokenHandler = handler.TokenHandle(jwtService)
	e.GET("/token", func(c echo.Context) error{
		token := tokenHandler.GetToken()
		if token != "" {
			return c.JSON(http.StatusOK, map[string]string{
				"token": token,
			})
		} else {
			return c.JSON(http.StatusUnauthorized, nil)
		}
	})
	e.GET("/employees", handler.GetAllEmployees)
	e.GET("/employees/:id", handler.GetEmployeeById)
	e.POST("/employees", handler.CreateEmployee)
	e.PUT("/employees/:id", handler.UpdateEmployee)
	e.DELETE("/employees/:id", handler.DeleteEmployee)

	return e
}
