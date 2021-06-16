package handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang-echo/dbconfig"
	"golang-echo/model"
	"golang-echo/service"
	"net/http"
	"strings"
	"time"
)

// CreateEmployee godoc
// @Summary create employee
// @Description create employee
// @Accept  json
// @Produce  json
// @Param employee body model.Employees true "Employee"
// @Success 200 {object} model.Employees
// @Router /employees [post]
func CreateEmployee(c echo.Context) error {
	req := new(model.Employees)
	if err := c.Bind(&req); err != nil {
		return err
	}

	dbconfig.Database.Create(&req)
	return c.JSON(http.StatusCreated, "Success create employee")
}

// GetAllEmployees godoc
// @Summary list employees
// @Description get all employees
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Employees
// @Router /employees [get]
func GetAllEmployees(c echo.Context) error  {
	auth := c.Request().Header.Get("Authorization")
	if !strings.Contains(auth, "Bearer") {

		return c.JSON(http.StatusBadRequest, "invalid token")
	}

	tokenStr := strings.Replace(auth, "Bearer ", "", -1)
	token, err := jwt.ParseWithClaims(tokenStr, &service.AuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.GetSecretKey()), nil
	})

	claims, ok := token.Claims.(*service.AuthCustomClaims)

	if !ok {
		return err
	}
	waktu := time.Now().UTC().Unix()

	if claims.ExpiresAt < waktu {
		return c.JSON(http.StatusBadRequest, "expired token")
	}
	e := new([]model.Employees)
	dbconfig.Database.Find(&e)

	return c.JSON(http.StatusCreated, e)
}

// GetEmployeeById godoc
// @Summary get employee by id
// @Description get by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Employee ID"
// @Success 200 {object} model.Employees
// @Router /employees/{id} [get]
func GetEmployeeById(c echo.Context) error {
	e := new(model.Employees)
	if err := dbconfig.Database.Where("id = ?", c.Param("id")).First(&e).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, e)
}

// UpdateEmployee godoc
// @Summary update employee
// @Description update employee
// @Accept  json
// @Produce  json
// @Param id path int true "Employee ID"
// @Param employee body model.Employees true "Employee"
// @Success 200 {object} model.Employees
// @Router /employees/{id} [put]
func UpdateEmployee(c echo.Context) error{
	e := new(model.Employees)
	req := new(model.Employees)

	if err := dbconfig.Database.Where("id = ?", c.Param("id")).First(&e).Error; err != nil {
		return err
	}

	if err := c.Bind(req); err != nil {
		return err
	}

	dbconfig.Database.Model(&e).Updates(&req)

	return c.JSON(http.StatusCreated, "employee updated")
}

// DeleteEmployee godoc
// @Summary delete employee
// @Description delete
// @Accept  json
// @Produce  json
// @Param id path int true "Employee ID"
// @Success 200 {object} model.Employees
// @Router /employees/{id} [delete]
func DeleteEmployee(c echo.Context) error {
	e := new(model.Employees)
	if err := dbconfig.Database.Where("id = ?", c.Param("id")).First(&e).Error; err != nil {
		return err
	}

	dbconfig.Database.Delete(&e)

	return c.JSON(http.StatusCreated, "employee deleted")
}



