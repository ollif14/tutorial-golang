package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Employees struct {
	Id int `json:"id"`
	Email_address string `json:"email_address"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "santosa2509"
	dbname   = "employees"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Successfully connected!")
	return db
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	rows, err := db.Query("SELECT * FROM employees")
	if err != nil {
		log.Fatal(err)
	}

	var employees []Employees 

	for rows.Next() {
		var employee Employees
		rows.Scan(&employee.Id, &employee.Email_address, &employee.First_name, & employee.Last_name)
		employees = append(employees, employee)
	}

	employeeBytes, _ := json.MarshalIndent(employees, "", "\t")
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(employeeBytes)

	defer rows.Close()
	defer db.Close()
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	var e Employees
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO employees (email_address, first_name, last_name) VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, e.Email_address, e.First_name, e.Last_name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

func main() {
	OpenConnection()
	http.HandleFunc("/", GetAllEmployees)
	http.HandleFunc("/insert", CreateEmployee)
	log.Fatal(http.ListenAndServe(":8080", nil))
}