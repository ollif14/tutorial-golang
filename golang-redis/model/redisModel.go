package model

type Student struct {
	Id int64 `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Gpa float64 `json:"gpa" binding:"required"`
}
