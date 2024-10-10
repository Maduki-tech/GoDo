package main

import (
	"github.com/Maduki-tech/GoDo/todo"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/add", addTodo)

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

	Todo := Todo.New()

}
