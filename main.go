package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vashu992/dating-app/controller.go"
	"github.com/vashu992/dating-app/middleware"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/user/create", controller.CreateUser).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/discover", middleware.AuthMiddleware(controller.Discover)).Methods("GET")
	r.HandleFunc("/swipe", middleware.AuthMiddleware(controller.Swipe)).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
