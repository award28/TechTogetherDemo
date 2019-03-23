package root

import (
	"net/http"
)

type Todo struct {
	ID       int    `json:"id"`
	Task     string `json:"task"`
	Finished bool   `json:"finished"`
}

type TodoRouter interface {
	UploadHandler(http.ResponseWriter, *http.Request) error
	UpdateHandler(http.ResponseWriter, *http.Request) error
	LookupHandler(http.ResponseWriter, *http.Request) error
}

type TodoService interface {
	Lookup(int) (*Todo, error)
	Update(*Todo) (*Todo, error)
	Upload(*Todo) *Todo
}

type TodoVault interface {
	GetByID(int) (*Todo, error)
	Modify(*Todo) (*Todo, error)
	Create(*Todo) *Todo
}
