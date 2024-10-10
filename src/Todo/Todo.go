package todo

type Todo struct {
	Todos []TodoElement
}

type TodoElement struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func New() Todo {
	return Todo{
		Todos: []TodoElement{},
	}
}

func (t *Todo) AddTodo(input TodoElement) {
	t.Todos = append(t.Todos, input)
}

func (t Todo) GetTodos() []TodoElement {
	return t.Todos
}
