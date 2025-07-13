package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/yusufborucu/go-projects/crud-api-jwt-gorm/config"
	"github.com/yusufborucu/go-projects/crud-api-jwt-gorm/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	log.Println("Server started at :8000")
	http.ListenAndServe(":8000", r)
}
