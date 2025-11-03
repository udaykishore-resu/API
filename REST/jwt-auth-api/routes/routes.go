package routes

import "github.com/gorilla/mux"

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Apply security headers to all routes
	r.Use(Securi)
}
