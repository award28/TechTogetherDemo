package main

import (
	"github.com/award28/TechTogetherDemo/internal/server"
	"github.com/award28/TechTogetherDemo/internal/todo"
)

func main() {
	todoVault := todo.NewTodoVault()
	todoService := todo.NewTodoService(todoVault)
	todoRouter := todo.NewTodoRouter(todoService, "/todos/")

	s := server.NewServer(todoRouter)
	s.Start()
}
