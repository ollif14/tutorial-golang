package controller

import (
	"api-with-gorm/config"
	"api-with-gorm/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllEmployees godoc
// @Summary list employees
// @Description get all employees
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Employees
// @Router /employees [get]
func GetAllEmployees(c *gin.Context) {
	var employees []model.Employees
	config.Database.Find(&employees)

	c.JSON(http.StatusOK, gin.H{"data": employees})
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
func GetEmployeeById(c *gin.Context) {
	var employee model.Employees
	if err := config.Database.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// CreateEmployee godoc
// @Summary create employee
// @Description create employee
// @Accept  json
// @Produce  json
// @Param employee body model.EmployeeRequest true "Employee"
// @Success 200 {object} model.Employees
// @Router /employees [post]
func CreateEmployee(c *gin.Context) {
	var request model.EmployeeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee := model.Employees{
		Email_address: request.Email_address,
		First_name:    request.First_name,
		Last_name:     request.Last_name}

	config.Database.Create(&employee)

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// UpdateEmployee godoc
// @Summary update employee
// @Description update employee
// @Accept  json
// @Produce  json
// @Param id path int true "Employee ID"
// @Param employee body model.EmployeeUpdateReq true "Employee"
// @Success 200 {object} model.Employees
// @Router /employees/{id} [put]
func UpdateEmployee(c *gin.Context) {

	var employee model.Employees
	if err := config.Database.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var request model.EmployeeUpdateReq
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in validate"})
		return
	}

	config.Database.Model(&employee).Updates(&model.Employees{
		Email_address: request.Email_address,
		First_name:    request.First_name,
		Last_name:     request.Last_name})

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// DeleteEmployee godoc
// @Summary delete employee
// @Description delete
// @Accept  json
// @Produce  json
// @Param id path int true "Employee ID"
// @Success 200 {object} model.Employees
// @Router /employees/{id} [delete]
func DeleteEmployee(c *gin.Context) {

	var employee model.Employees

	if err := config.Database.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.Database.Delete(&employee)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
