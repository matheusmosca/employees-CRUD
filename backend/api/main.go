package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// ctx := context.Background()
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	dbConfig := "host=localhost user=postgres port=5493 dbname=nutcacheDB password=postgres sslmode=disable"

	sqlDB, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println(sqlDB)
}
