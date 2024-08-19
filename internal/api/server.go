package api

import (
	"github.com/kama-meshi/api-first-example-go/internal/service"
	api "github.com/kama-meshi/api-first-example-go/pkg"
	"github.com/labstack/echo/v4"
)

type Server struct {
	service *service.TodoService
}

var _ api.ServerInterface = &Server{}

func NewServer() *Server {
	return &Server{
		service: service.NewTodoService(),
	}
}

// GetTodos implements api.ServerInterface.
func (s *Server) GetTodos(ctx echo.Context) error {
	todos := s.service.GetTodos()
	ctx.JSON(200, todos)
	return nil
}

// GetTodosTodoId implements api.ServerInterface.
func (s *Server) GetTodosTodoId(ctx echo.Context, todoId int) error {
	todo := s.service.GetTodoById(todoId)
	if todo != nil {
		ctx.JSON(200, todo)
		return nil
	} else {
		return echo.NewHTTPError(404, "Not Found")
	}
}

// PostTodos implements api.ServerInterface.
func (s *Server) PostTodos(ctx echo.Context) error {
	var todoInput api.TodoInput
	if err := ctx.Bind(&todoInput); err != nil {
		return echo.NewHTTPError(400, "Invalid request payload")
	}

	s.service.CreateTodo(&todoInput)
	ctx.JSON(201, nil)
	return nil
}

// PutTodosTodoId implements api.ServerInterface.
func (s *Server) PutTodosTodoId(ctx echo.Context, todoId int) error {
	var todoInput api.TodoInput
	if err := ctx.Bind(&todoInput); err != nil {
		return echo.NewHTTPError(400, "Invalid request")
	}
	s.service.UpdateTodoById(todoId, &todoInput)
	return nil
}

// DeleteTodosTodoId implements api.ServerInterface.
func (s *Server) DeleteTodosTodoId(ctx echo.Context, todoId int) error {
	s.service.DeleteTodoById(todoId)
	return nil
}
