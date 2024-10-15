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

	fmt.Println("Starting main function")

	router := mux.NewRouter()

	todos = append(todos, 1)
	todos = append(todos, 2)
	todos = append(todos, 3)

	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos", addTodos).Methods("POST")
	router.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")
	router.HandleFunc("/todos/{id}", getTodoById).Methods("GET")

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func getTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Implementing getTodos")
	//NewEncoder(w) , which will write the JSON output to the response writer w .
	//The person struct was encoded to JSON and written to the response using encoder
	json.NewEncoder(w).Encode(todos)
	// w.WriteHeader(200)
}

func addTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Implementing create todo")

	todos = append(todos, len(todos)+1)
	w.WriteHeader(200)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {

	incomingId := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(incomingId)
	fmt.Printf("Incoming Id to delete %v", id)

	if id >= len(todos) {
		w.Write([]byte("Invalid Id"))
		return
	}

	todos = append(todos[:id], todos[id+1:]...)
	fmt.Printf("Updated todos %v", todos)
	w.WriteHeader(200)
}

func getTodoById(w http.ResponseWriter, r *http.Request) {

	incomingId := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(incomingId)
	fmt.Printf("Incoming Id to delete %v", id)

	if id < 1 || id >= len(todos)+1 {
		w.Write([]byte("Invalid Id"))
		return
	}

	json.NewEncoder(w).Encode(todos[id-1])
	// w.WriteHeader(200)
}
