package controller

import (
	"github.com/labstack/echo/v4"
	"golang-mongodb/dbconfiguration"
	"golang-mongodb/model"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)
const database = "testDB"
const collection = "table1"
var session = dbconfiguration.ConnectToMongo()

// CreateEmployee godoc
// @Summary create employee
// @Description create employee
// @Accept  json
// @Produce  json
// @Param employee body model.Employees true "Employee"
// @Success 200 {object} model.Employees
// @Router /employees [post]
func CreateEmployee(c echo.Context) error {
	var req model.Employee
	req.Id = bson.NewObjectId()
	if err := c.Bind(&req); err != nil {
		return err
	}
	test := session.DB(database).C(collection)
	if err := test.Insert(&req); err != nil{
		return c.JSON(http.StatusCreated, err)
	}
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
	var e []model.Employee
	test := session.DB(database).C(collection)
	if err := test.Find(nil).All(&e); err != nil{
		return c.JSON(http.StatusCreated, err)
	}
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
	eId := c.Param("id")
	var e model.Employee
	test := session.DB(database).C(collection)
	if err := test.Find(bson.M{"_id" : bson.ObjectIdHex(eId)}).One(&e); err != nil{
		return c.JSON(http.StatusCreated, "not found")
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
	var exist model.Employee
	var req model.Employee
	eId := c.Param("id")

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusCreated, err)
	}

	test := session.DB(database).C(collection)
	if err := test.Find(bson.M{"_id" : bson.ObjectIdHex(eId)}).One(&exist); err != nil{
		return c.JSON(http.StatusCreated, err)
	}

	req.Id = exist.Id
	if err := test.Update(&exist, &req); err != nil{
		return c.JSON(http.StatusCreated, err)
	}
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
	var eDel model.Employee
	eId := c.Param("id")
	test := session.DB(database).C(collection)
	if err := test.Find(bson.M{"_id" : bson.ObjectIdHex(eId)}).One(&eDel); err != nil{
		return c.JSON(http.StatusCreated, err)
	}

	if err := test.Remove(&eDel); err != nil{
		return c.JSON(http.StatusCreated, err)
	}
	return c.JSON(http.StatusCreated, "employee deleted")
}



