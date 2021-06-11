package service

import (
	"github.com/h-u-m-a-n/todo-app"
	"github.com/h-u-m-a-n/todo-app/package/repository"
)

type TodoItemService struct{
	repo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil{
		// list doesn't exitst or doesn't belong to user
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error){
	return s.repo.GetAll(userId, listId)
}