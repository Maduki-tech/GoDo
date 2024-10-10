package todo_test

import (
	todo "github.com/Maduki-tech/GoDo/src/Todo"
	"testing"
)

func NewTodo(t *testing.T) {
	todo := todo.New()
	if todo.Todos == nil {
		t.Error("Expected todo.Todos to be initialized")
	}
}

func AddTodo_Test(t *testing.T) {
	todo := todo.New()
	todo.AddTodo("Test")
	if len(todo.Todos) != 1 {
		t.Error("Expected todo.Todos to have 1 element")
	}
}

func GetTodos_Test(t *testing.T) {
	todo := todo.New()
	todo.AddTodo("Test")
	todos := todo.GetTodos()
	if len(todos) != 1 {
		t.Error("Expected todo.GetTodos() to return 1 element")
	}
}
