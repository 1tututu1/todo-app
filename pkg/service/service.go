package service

import (
	"github.com/1tututu1/pkg/repository"
	todoapp "github.com/1tututu1"
)

type Authorization interface {
	CreateUser(user todoapp.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
