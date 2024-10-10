package main

import (
	"github.com/Maduki-tech/GoDo/src/Todo"
	"log"
	"net/http"
)

var Todo = todo.New()

func main() {
	http.HandleFunc("/add", addTodo)
	http.HandleFunc("/getAll", getAllTodos)

	server := &http.Server{
		Addr: ":8080",
	}

	log.Fatal(server.ListenAndServe())
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	Todo.AddTodo("New Todo")
}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	todos := Todo.GetTodos()

	for _, todo := range todos {
		log.Println(todos)
		w.Write([]byte(todo.Title))
	}

	w.WriteHeader(http.StatusOK)

}
