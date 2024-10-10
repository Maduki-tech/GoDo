package todo

type Todo struct {
	Todos []TodoElement
}

type TodoElement struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t Todo) New() Todo {

	return Todo{
		Todos: make([]TodoElement, 10),
	}

}

func (t Todo) AddTodo(input Todo) {

}
