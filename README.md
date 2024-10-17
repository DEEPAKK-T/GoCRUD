# GoCRUD


GoCRUD is a simple Go application designed for beginners to learn how to build a backend API for managing a Todo list. Users can perform Create, Read, Update, and Delete (CRUD) operations on Todo items.

## Features

No database required.
Simple and beginner-friendly.
Full CRUD functionality on Todo items.

## Prerequisites

Go (version 1.xx or higher)

## Usage
You can interact with the API using tools like curl or Postman.

### Endpoints:

1. Create a Todo
    - POST /todos

2. Get all Todos
    - GET /todos

3. Get Todo by Id
    - GET /todos{id}

4. Update Todo by Id
    - PUT /todos{id}
      Body :
      {
        "newValue" : 30
      }

5. Delete Todo
    -  DELETE /todos{id}