package repository

import (
	"golang-redis/config"
	"golang-redis/model"
)

func Save(req *model.Student) error{
	if err := config.Client.SetKey("student",req, 0); err != nil{
		return err
	}
	return nil
}

func GetBy(key string) (*model.Student, error) {
	var student model.Student
	if err := config.Client.GetKey(key, &student); err != nil {
		return nil, err
	}
	return &student, nil
}

func Delete(key string) error {
	err := config.Client.C.Del(key)
	if err != nil {
		return nil
	}
	return nil
}