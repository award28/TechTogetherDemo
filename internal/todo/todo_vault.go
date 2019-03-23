package todo

import (
	"errors"
	"github.com/award28/TechTogetherDemo/internal/root"
)

type TodoMap = map[int]*root.Todo

type TodoVault struct {
	todos TodoMap
	next  int
}

func NewTodoVault() *TodoVault {
	return &TodoVault{
		todos: make(TodoMap),
		next:  1,
	}
}

func (tv *TodoVault) GetByID(ID int) (*root.Todo, error) {
	if todo, ok := tv.todos[ID]; !ok {
		return nil, errors.New("Todo does not exist.")
	} else {
		return todo, nil
	}
}

func (tv *TodoVault) Modify(todo *root.Todo) (*root.Todo, error) {
	if _, ok := tv.todos[todo.ID]; !ok {
		return nil, errors.New("Todo does not exist.")
	}
	tv.todos[todo.ID] = todo
	return todo, nil
}

func (tv *TodoVault) Create(todo *root.Todo) *root.Todo {
	todo.ID = tv.next
	tv.todos[tv.next] = todo
	tv.next++
	return todo
}
