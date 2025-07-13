package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yusufborucu/go-projects/crud-api-jwt-gorm/config"
	"github.com/yusufborucu/go-projects/crud-api-jwt-gorm/models"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	var todos []models.Todo
	config.DB.Where("user_id = ?", userID).Find(&todos)
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todo.UserID = userID

	config.DB.Create(&todo)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	var todo models.Todo
	result := config.DB.First(&todo, id)

	if result.Error != nil || todo.UserID != userID {
		http.Error(w, "Todo not found or unauthorized", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&todo)
	config.DB.Save(&todo)
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	var todo models.Todo
	result := config.DB.First(&todo, id)

	if result.Error != nil || todo.UserID != userID {
		http.Error(w, "Todo not found or unauthorized", http.StatusNotFound)
		return
	}

	config.DB.Delete(&todo)
	w.WriteHeader(http.StatusNoContent)
}
