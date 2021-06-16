package dbconfig

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectToDB() {
	dsn := "host=localhost user=postgres password=santosa2509 dbname=employees port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	Database = db

	fmt.Println("Success to Connect")
}
