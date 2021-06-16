package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"golang-mongo-convert-csv/config"
	"golang-mongo-convert-csv/model"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func main(){
	config.ConnectToMongo()
	GetDataAndConvertToCSV()
}

func GetDataAndConvertToCSV() error{
	var records []string
	data := &model.Employee{}

	result, err := config.Employeedb.Find(context.Background(), bson.M{})
	if err != nil{
		panic(err)
	}

	csvFile, err := os.Create("employee.csv")
	if err != nil {
		panic(err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for result.Next(context.Background()){
		if err := result.Decode(data); err != nil{
			panic(err)
		}

		out, err := json.Marshal(data)
		if err != nil {
			panic (err)
		}
		csvwriter.Write(append(records, string(out)))
	}

	csvwriter.Flush()
	csvFile.Close()
	return nil
}
