package users

import "context"

type Repository interface {
	Create(ctx context.Context, user *User) error
	List() ([]User, error)
	GetByID(id string) (*User, error)
	Delete(ctx context.Context, id string) error
}
