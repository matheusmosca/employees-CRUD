package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusmosca/nutcache/api/handler"
	"github.com/matheusmosca/nutcache/infrastructure/repo"
	"github.com/matheusmosca/nutcache/model"
	"github.com/matheusmosca/nutcache/service"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbConfig := "host=localhost user=postgres port=5493 dbname=nutcacheDB password=postgres sslmode=disable"

	DB, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.Employee{})

	r := mux.NewRouter()
	employeeRepo := repo.NewEmployeeRepo(DB)
	employeeService := service.NewEmployeeService(employeeRepo)

	handler.MakeEmployeeEndpoints(r, employeeService)

	fmt.Println("Listening in 3030...")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	http.ListenAndServe(":3030", c.Handler(r))
}
