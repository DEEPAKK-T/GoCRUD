package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Backend API to manage Todo items")
	router := mux.NewRouter()

	//CRUD
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos", createTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", getTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", updateTodo).Methods("UPDATE")
	router.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")

	fmt.Println("Server listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("implement")
	w.WriteHeader(200)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
