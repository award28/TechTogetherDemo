// +build unit

package todo_test

import (
	"github.com/award28/TechTogetherDemo/internal/root"
	"github.com/award28/TechTogetherDemo/internal/todo"
	"testing"
)

func Test_TodoService(t *testing.T) {
	todoService := todo.NewTodoService(
		todo.NewTodoVault(),
	)

	actualTodo := &root.Todo{
		ID:       -1,
		Task:     "Hack",
		Finished: false,
	}
	t.Run("Upload todo", upload_todo(todoService, actualTodo))
	t.Run("Lookup todo", lookup_todo(todoService, actualTodo.Task))
	actualTodo.Finished = true
	t.Run("Update todo", update_todo(todoService, actualTodo))
}

func upload_todo(todoService root.TodoService, actualTodo *root.Todo) func(*testing.T) {
	return func(t *testing.T) {
		res_todo := todoService.Upload(actualTodo)
		actualTodo.ID = res_todo.ID
		if res_todo != actualTodo {
			t.Errorf("Mismatched.")
		}
	}
}

func lookup_todo(todoService root.TodoService, task string) func(*testing.T) {
	return func(t *testing.T) {
		responseTodo, err := todoService.Lookup(1)
		if err != nil {
			t.Errorf("Unable to lookup todo: %s", err)
		}

		if responseTodo.Task != task {
			t.Errorf("Response Task: `%s` Expected Task: %s",
				responseTodo.Task, task)
		}
	}
}

func update_todo(todoService root.TodoService, actualTodo *root.Todo) func(*testing.T) {
	return func(t *testing.T) {
		responseTodo, err := todoService.Update(actualTodo)
		if err != nil {
			t.Errorf("Unable to lookup todo: %s", err)
		}

		if responseTodo.Finished != actualTodo.Finished {
			t.Errorf("Response finished: `%t` Expected finished: %t",
				responseTodo.Finished, actualTodo.Finished)
		}
	}
}
