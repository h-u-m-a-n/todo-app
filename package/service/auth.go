package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/h-u-m-a-n/todo-app"
	"github.com/h-u-m-a-n/todo-app/package/repository"
)

const salt = "hdsifhi2o13y473"


type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}