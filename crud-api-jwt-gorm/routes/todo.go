package routes

import (
	"github.com/gorilla/mux"

	"github.com/yusufborucu/go-projects/crud-api-jwt-gorm/controllers"
	"github.com/yusufborucu/go-projects/crud-api-jwt-gorm/middleware"
)

func RegisterTodoRoutes(r *mux.Router) {
	s := r.PathPrefix("/todos").Subrouter()
	s.Use(middleware.JWTMiddleware)

	s.HandleFunc("", controllers.GetTodos).Methods("GET")
	s.HandleFunc("", controllers.CreateTodo).Methods("POST")
	s.HandleFunc("/{id}", controllers.UpdateTodo).Methods("PUT")
	s.HandleFunc("/{id}", controllers.DeleteTodo).Methods("DELETE")
}
