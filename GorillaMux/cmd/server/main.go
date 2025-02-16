package main

import (
	"gorillamux/internal/handlers"
	"gorillamux/internal/routes"
	"gorillamux/internal/storage"
	"log"
	"net/http"
)

func main() {
	// MySQL connection string format: "user:password@tcp(host:port)/dbname"
	storage, err := storage.NewEmployeeStorage("root:Vanshika@9@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	employeeHandler := handlers.NewEmployeeHandler(storage)
	router := routes.SetupRouter(employeeHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
