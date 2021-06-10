package service

import (
	"github.com/h-u-m-a-n/todo-app"
	"github.com/h-u-m-a-n/todo-app/package/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error){
	return s.repo.Create(userId, list) 
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error){
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetListByID(userId, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) DeleteListByID(userId, listId int) error{
	return s.repo.DeleteListByID(userId, listId)
}

func (s *TodoListService) UpdateListById(userId, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil{
		return err
	}
	return s.repo.UpdateListById(userId, listId, input) 
}