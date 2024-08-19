package service

import (
	"time"

	api "github.com/kama-meshi/api-first-example-go/pkg"
)

// TODO Service
type TodoService struct {
	// TODO item repository
	todos []api.Todo
}

type Option func([]api.Todo) []api.Todo

func NewTodoService() *TodoService {
	return &TodoService{
		todos: make([]api.Todo, 0),
	}
}

// Get all TODO items
func (h *TodoService) GetTodos() []api.Todo {
	todos := h.todos
	return todos
}

// Get TODO item
func (h *TodoService) GetTodoById(id int) *api.Todo {
	for _, todo := range h.todos {
		if todo.Id == id {
			return &todo
		}
	}
	return nil
}

// Create new TODO item
func (h *TodoService) CreateTodo(todo *api.TodoInput) {
	newTodo := &api.Todo{
		Id:          len(h.todos) + 1,
		Title:       todo.Title,
		Completed:   todo.Completed,
		Description: todo.Description,
		CreatedAt:   time.Now(),
	}
	h.todos = append(h.todos, *newTodo)
}

// Update TODO item
func (h *TodoService) UpdateTodoById(id int, todo *api.TodoInput) {
	for i, t := range h.todos {
		if t.Id == id {
			h.todos[i].Title = todo.Title
			h.todos[i].Completed = todo.Completed
			h.todos[i].Description = todo.Description
			return
		}
	}
}

// Delete TODO item
func (h *TodoService) DeleteTodoById(id int) {
	for i, t := range h.todos {
		if t.Id == id {
			h.todos = append(h.todos[:i], h.todos[i+1:]...)
			return
		}
	}
}
