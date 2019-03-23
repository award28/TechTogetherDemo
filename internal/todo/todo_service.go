package todo

import (
	"github.com/award28/TechTogetherDemo/internal/root"
)

type TodoService struct {
	todoVault root.TodoVault
}

func NewTodoService(iv root.TodoVault) *TodoService {
	return &TodoService{iv}
}

func (is *TodoService) Upload(todo *root.Todo) *root.Todo {
	return is.todoVault.Create(todo)
}

func (is *TodoService) Update(todo *root.Todo) (*root.Todo, error) {
	return is.todoVault.Modify(todo)
}

func (is *TodoService) Lookup(ID int) (*root.Todo, error) {
	return is.todoVault.GetByID(ID)
}
