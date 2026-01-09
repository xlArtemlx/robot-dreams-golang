package users

import (
	"context"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, id string, name string) (*User, error) {
	user := &User{ID: id, Name: name}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) ListUsers() ([]User, error) {
	return s.repo.List()
}

func (s *Service) GetUser(userID string) (*User, error) {
	return s.repo.GetByID(userID)
}

func (s *Service) DeleteUser(ctx context.Context, userID string) error {
	return s.repo.Delete(ctx, userID)
}
