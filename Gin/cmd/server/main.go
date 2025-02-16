package main

import (
	"gin-api/internal/controllers"
	"gin-api/internal/routes"
	"gin-api/internal/storage"
	"log"
)

func main() {
	// MySQL connection string format: "user:password@tcp(host:port)/dbname"
	storage, err := storage.NewEmployeeStorage("root:password@tcp(localhost:3306)/employeedb")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	employeeCtrl := controllers.NewEmployeeController(storage)
	router := routes.SetupRouter(employeeCtrl)

	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
