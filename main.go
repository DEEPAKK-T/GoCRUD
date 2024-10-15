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

// struct to get the json request body
type UpdateRequest struct {
	NewValue int `json:"newValue"`
}

func main() {

	fmt.Println("Starting main function")

	//Initializing the router
	router := mux.NewRouter()

	//Intial todo list
	todos = append(todos, 1)
	todos = append(todos, 2)
	todos = append(todos, 3)

	//CRUD APIs with HTTP methods.
	router.HandleFunc("/todos", getTodos).Methods("GET")            //Read all todos
	router.HandleFunc("/todos", addTodos).Methods("POST")           //Create a new todo
	router.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")  //Delete a specific todo by ID
	router.HandleFunc("/todos/{id}", getTodoById).Methods("GET")    //Read a specific todo by ID
	router.HandleFunc("/todos/{id}", updateTodoById).Methods("PUT") //Update a specific todo by ID

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func getTodos(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Implementing getTodos")
	//NewEncoder(w) , which will write the JSON output to the response writer w .
	//The person struct was encoded to JSON and written to the response using encoder
	json.NewEncoder(w).Encode(todos)
}

func addTodos(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Implementing create todo")
	todos = append(todos, len(todos)+1)
	w.WriteHeader(200)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {

	//do initial check of array
	if len(todos) == 0 {
		w.Write([]byte("The to-do list is already empty, so there's not much left to delete."))
		return
	}

	//get id as per human understanding
	incomingId := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(incomingId)
	fmt.Printf("Incoming Id to delete %v", id)

	if id > len(todos) || id <= 0 {
		w.Write([]byte("Invalid Id"))
		return
	}
	//decrement id as per array indexing
	id = id - 1

	todos = append(todos[:id], todos[id+1:]...)
	fmt.Printf("Updated todos %v", todos)
	w.WriteHeader(200)
}

func getTodoById(w http.ResponseWriter, r *http.Request) {

	//do initial check of array
	if len(todos) == 0 {
		w.Write([]byte("The to-do list is already empty, so there's not much left to get."))
		http.Error(w, "The to-do list is already empty, so there's not much left to get.", http.StatusBadRequest)
		return
	}

	incomingId := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(incomingId)
	fmt.Printf("Incoming Id to delete %v", id)

	if id < 1 || id >= len(todos)+1 {
		w.Write([]byte("Invalid Id"))
		return
	}

	json.NewEncoder(w).Encode(todos[id-1])
}

func updateTodoById(w http.ResponseWriter, r *http.Request) {

	// Work sequence:
	// 1. Retrieve the ID from the path parameters.
	// 2. Parse the request body to get the updated value.
	// 3. If the todo list is empty, reject the request.
	// 4. Check if the ID exists within the current todo list (based on array length).
	// 5. Ensure the updated value doesn't already exist in the list to avoid duplicates.

	if len(todos) == 0 {
		w.Write([]byte("The to-do list is already empty, so there's not much left to update."))
		http.Error(w, "The to-do list is already empty, so there's not much left to update.", http.StatusBadRequest)
		return
	}

	//Read the id param
	incomingId := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(incomingId)
	fmt.Printf("\nIncoming Id to update %v", id)

	if id <= 0 || id > len(todos) {
		w.Write([]byte("Invalid Id"))
		return
	}
	//decrement id as per array indexing
	id = id - 1

	var requestBody UpdateRequest
	//Read the request body to get the updated value.
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newValue := requestBody.NewValue

	fmt.Printf("\n incoming updated value %v", newValue)

	if contains(todos, newValue) {
		fmt.Printf("%d already in todos list, won't accept duplicate values in the list", newValue)
		http.Error(w, fmt.Sprintf("%d already in todos list, won't accept duplicate values in the list", newValue), http.StatusBadRequest)
		return
	}

	if !updateElementInTodo(todos, id, newValue) {
		fmt.Println("Fail to update element in todo list. Internal server error")
		http.Error(w, "Fail to update element in todo list. Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Printf("Updated todo list is %v", todos)
	w.WriteHeader(200)

}

func contains(arr []int, value int) bool {

	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false

}

func updateElementInTodo(arr []int, incomingId int, newValue int) bool {

	for id, _ := range arr {
		if id == incomingId {
			arr[id] = newValue
			return true
		}
	}

	return false
}
