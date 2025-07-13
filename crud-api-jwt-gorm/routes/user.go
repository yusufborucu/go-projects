package routes

import (
	"github.com/gorilla/mux"

	"github.com/yusufborucu/go-projects/crud-api-jwt-gorm/controllers"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/signup", controllers.Signup).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	RegisterTodoRoutes(r)
}
