package handlers

import (
	"encoding/json"
	"gorillamux/internal/models"
	"gorillamux/internal/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EmployeeHandler struct {
	storage *storage.EmployeeStorage
}

func NewEmployeeHandler(storage *storage.EmployeeStorage) *EmployeeHandler {
	return &EmployeeHandler{storage: storage}
}

func (h *EmployeeHandler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.storage.GetEmployees()
	if err != nil {
		return
	}
	responseJson(w, http.StatusOK, employees)
}

func (h *EmployeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responseJson(w, http.StatusBadRequest, "invalid employee ID")
		return
	}

	employee, err := h.storage.GetEmployee(id)
	if err != nil {
		return
	}
	responseJson(w, http.StatusOK, employee)
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		responseJson(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	if err := h.storage.CreateEmployee(&emp); err != nil {
		responseJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJson(w, http.StatusCreated, emp)
}

func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responseJson(w, http.StatusBadRequest, "invalid employee ID")
		return
	}

	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		responseJson(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	if err := h.storage.UpdateEmployee(id, &emp); err != nil {
		responseJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	emp.ID = id
	responseJson(w, http.StatusOK, emp)
}

func (h *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responseJson(w, http.StatusBadRequest, "invalid employee ID")
		return
	}

	if err := h.storage.DeleteEmployee(id); err != nil {
		responseJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Helper functions remain the same
func responseJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
