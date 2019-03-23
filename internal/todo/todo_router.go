package todo

import (
	"github.com/MercuryThePlanet/rest-tools"
	"github.com/award28/TechTogetherDemo/internal/root"
	"net/http"
)

type TodoRouter struct {
	todoService root.TodoService
	RelPath     string
}

func NewTodoRouter(TodoService root.TodoService, relPath string) *TodoRouter {
	return &TodoRouter{
		todoService: TodoService,
		RelPath:     relPath,
	}
}

func (ir *TodoRouter) InitRoutes() {
	http.Handle(ir.RelPath, tools.NewRestHelper(tools.MethodMap{
		http.MethodGet:  ir.LookupHandler,
		http.MethodPost: ir.UploadHandler,
		http.MethodPut:  ir.UpdateHandler,
	}).JsonErrHandler())
}

func (ir *TodoRouter) UploadHandler(w http.ResponseWriter, r *http.Request) error {
	var new_todo root.Todo
	err := tools.Unmarshal(&new_todo, r)
	if err != nil {
		return tools.StatusError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	resTodo := ir.todoService.Upload(&new_todo)
	return tools.ServeJsonRes(w, http.StatusCreated, resTodo)
}

func (ir *TodoRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := tools.PathParamToInt(r.URL.Path, ir.RelPath)
	if err != nil {
		return tools.StatusError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	var new_todo root.Todo
	err = tools.Unmarshal(&new_todo, r)
	if err != nil {
		return tools.StatusError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}
	new_todo.ID = id

	resTodo, err := ir.todoService.Update(&new_todo)
	if err != nil {
		return tools.StatusError{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}
	return tools.ServeJsonRes(w, http.StatusCreated, resTodo)
}

func (ir *TodoRouter) LookupHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := tools.PathParamToInt(r.URL.Path, ir.RelPath)
	if err != nil {
		return tools.StatusError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	todo, err := ir.todoService.Lookup(id)
	if err != nil {
		return tools.StatusError{
			Code: http.StatusNotFound,
			Err:  err,
		}
	}
	return tools.ServeJsonRes(w, http.StatusOK, todo)
}
