package routes

import (
	"gorillamux/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter(employeeHandler *handlers.EmployeeHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/employees", employeeHandler.GetEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", employeeHandler.GetEmployee).Methods("GET")
	router.HandleFunc("/employees", employeeHandler.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", employeeHandler.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employees/{id}", employeeHandler.DeleteEmployee).Methods("DELETE")

	return router
}
