// +build unit

package todo_test

import (
	"bytes"
	"github.com/award28/TechTogetherDemo/internal/todo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_TodoRouter(t *testing.T) {
	todoService := todo.NewTodoService(todo.NewTodoVault())
	todoRouter := todo.NewTodoRouter(todoService, "/todos/")

	t.Run("Test Todo Router Upload", uploadTest(todoRouter))
	t.Run("Test Todo Router Lookup", lookupTest(todoRouter))
	t.Run("Test Todo Router Update", updateTest(todoRouter))
}

func uploadTest(ir *todo.TodoRouter) func(*testing.T) {
	return func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
                                "id": 1,
                                "task": "Hacking",
                                "finished": false
			     }`))

		req, err := http.NewRequest("POST", ir.RelPath, body)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(wrapper(ir.UploadHandler, t))
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}
	}
}

func lookupTest(ir *todo.TodoRouter) func(*testing.T) {
	return func(t *testing.T) {
		req, err := http.NewRequest("GET", ir.RelPath+"1", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(wrapper(ir.LookupHandler, t))
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	}
}

func updateTest(ir *todo.TodoRouter) func(*testing.T) {
	return func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
                                "id": 1,
                                "task": "Hacking",
                                "finished": false
			     }`))

		req, err := http.NewRequest("PUT", ir.RelPath+"1", body)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(wrapper(ir.LookupHandler, t))
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	}
}

type customHandler = func(http.ResponseWriter, *http.Request) error
type handlerFunc = func(http.ResponseWriter, *http.Request)

func wrapper(f customHandler, t *testing.T) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			t.Error(err)
		}
	}
}
