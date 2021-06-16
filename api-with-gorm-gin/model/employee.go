package model

type Employees struct {
	Id            int    `json:"id" gorm:"primary_key"`
	Email_address string `json:"email_address"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
}

//for create
type EmployeeRequest struct {
	Email_address string `json:"email_address" binding:"required"`
	First_name    string `json:"first_name" binding:"required"`
	Last_name     string `json:"last_name" binding:"required"`
}

//for update
type EmployeeUpdateReq struct {
	Email_address string `json:"email_address"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
}
