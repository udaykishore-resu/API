package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	username string `json:"username"`
	password string `json:"password"`
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("payload error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", postLogin).Methods("POST")
	log.Println("Server starting at port 8080")
	log.Println(http.ListenAndServe(":8080", router))
}
