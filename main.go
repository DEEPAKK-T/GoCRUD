package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var todos []int

func main() {
	fmt.Println("Backend API to manage Todo items")
	router := mux.NewRouter()

	todos = append(todos, 1)
	todos = append(todos, 2)

	fmt.Printf("\n todos %v", todos)

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
	json.NewEncoder(w).Encode(todos)
	w.WriteHeader(200)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	todos = append(todos, len(todos)+1)
	fmt.Printf("\nTODO list %v", todos)
	w.WriteHeader(200)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	fmt.Printf("\n Id is %v", id)
	w.Write([]byte(id))
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	if id > len(todos)-1 {
		w.Write([]byte("Invalid ID"))
		return
	}
	todos = append(todos[:id], todos[id+1:]...)
	fmt.Printf("\nTODO list %v", todos)
	w.WriteHeader(200)
}
