package documentstore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStore_CreateGetDeleteCollection(t *testing.T) {
	ctx := context.Background()
	st := NewStore()

	col, err := st.CreateCollection(ctx, "users", &CollectionConfig{PrimaryKey: "id"})
	require.NoError(t, err)
	require.NotNil(t, col)

	col2, err := st.GetCollection("users")
	require.NoError(t, err)
	require.Equal(t, col, col2)

	err = st.DeleteCollection(ctx, "users")
	require.NoError(t, err)

	_, err = st.GetCollection("users")
	require.ErrorIs(t, err, ErrCollectionNotFound)
}

func TestStore_CreateCollection_AlreadyExists(t *testing.T) {
	ctx := context.Background()
	st := NewStore()

	_, err := st.CreateCollection(ctx, "users", &CollectionConfig{PrimaryKey: "id"})
	require.NoError(t, err)

	_, err = st.CreateCollection(ctx, "users", &CollectionConfig{PrimaryKey: "id"})
	require.ErrorIs(t, err, ErrCollectionAlreadyExists)
}

func TestStore_CreateCollection_Validation(t *testing.T) {
	ctx := context.Background()
	st := NewStore()

	_, err := st.CreateCollection(ctx, "", &CollectionConfig{PrimaryKey: "id"})
	require.ErrorIs(t, err, ErrEmptyCollectionName)

	_, err = st.CreateCollection(ctx, "users", nil)
	require.ErrorIs(t, err, ErrNilCollectionConfig)
}
