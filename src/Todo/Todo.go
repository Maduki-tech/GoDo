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

func (t *Todo) AddTodo(title string) {
	newID := len(t.Todos) + 1
	todo := TodoElement{
		ID:        newID,
		Title:     title,
		Completed: false,
	}

	t.Todos = append(t.Todos, todo)
}

func (t Todo) GetTodos() []TodoElement {
	return t.Todos
}
