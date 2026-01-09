package users_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/xlArtemlx/robot-dreams-golang/users"
	"github.com/xlArtemlx/robot-dreams-golang/users/mocks"
)

func TestService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepository(ctrl)
	service := users.NewService(repo)

	ctx := context.Background()

	repo.
		EXPECT().
		Create(ctx, &users.User{
			ID:   "111",
			Name: "Obi_Wan",
		}).
		Return(nil)

	user, err := service.CreateUser(ctx, "111", "Obi_Wan")

	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "111", user.ID)
	require.Equal(t, "Obi_Wan", user.Name)
}

func TestService_GetUser_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepository(ctrl)
	service := users.NewService(repo)

	repo.
		EXPECT().
		GetByID("404").
		Return(nil, users.ErrUserNotFound)

	user, err := service.GetUser("404")

	require.Nil(t, user)
	require.ErrorIs(t, err, users.ErrUserNotFound)
}
