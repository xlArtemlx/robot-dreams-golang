package users

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockUserRepo struct {
	createFn func(ctx context.Context, user *User) error
	getFn    func(id string) (*User, error)
	listFn   func() ([]User, error)
	deleteFn func(ctx context.Context, id string) error
}

func (m *mockUserRepo) Create(ctx context.Context, user *User) error {
	return m.createFn(ctx, user)
}

func (m *mockUserRepo) GetByID(id string) (*User, error) {
	return m.getFn(id)
}

func (m *mockUserRepo) List() ([]User, error) {
	return m.listFn()
}

func (m *mockUserRepo) Delete(ctx context.Context, id string) error {
	return m.deleteFn(ctx, id)
}

func TestService_CreateUser(t *testing.T) {
	ctx := context.Background()

	repo := &mockUserRepo{
		createFn: func(ctx context.Context, user *User) error {
			require.Equal(t, "111", user.ID)
			require.Equal(t, "Obi_Wan", user.Name)
			return nil
		},
	}

	service := NewService(repo)

	user, err := service.CreateUser(ctx, "111", "Obi_Wan")

	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "111", user.ID)
	require.Equal(t, "Obi_Wan", user.Name)
}

func TestService_GetUser_NotFound(t *testing.T) {

	repo := &mockUserRepo{
		getFn: func(id string) (*User, error) {
			return nil, ErrUserNotFound
		},
	}

	service := NewService(repo)

	user, err := service.GetUser("404")

	require.Nil(t, user)
	require.ErrorIs(t, err, ErrUserNotFound)
}
