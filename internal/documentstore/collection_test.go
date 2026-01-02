package documentstore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCollection_PutGetDelete(t *testing.T) {
	ctx := context.Background()
	st := NewStore()

	col, err := st.CreateCollection(ctx, "users", &CollectionConfig{PrimaryKey: "id"})
	require.NoError(t, err)

	doc := Document{
		Fields: map[string]DocumentField{
			"id":   {Type: DocumentFieldTypeString, Value: "222"},
			"name": {Type: DocumentFieldTypeString, Value: "luke_skywalker"},
		},
	}

	err = col.Put(ctx, doc)
	require.NoError(t, err)

	got, err := col.Get("222")
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "luke_skywalker", got.Fields["name"].Value)

	err = col.Delete(ctx, "222")
	require.NoError(t, err)

	_, err = col.Get("222")
	require.ErrorIs(t, err, ErrDocumentNotFound)
}

func TestCollection_Put_PrimaryKeyValidation(t *testing.T) {
	ctx := context.Background()
	st := NewStore()

	col, err := st.CreateCollection(ctx, "users", &CollectionConfig{PrimaryKey: "id"})
	require.NoError(t, err)

	err = col.Put(ctx, Document{Fields: map[string]DocumentField{
		"name": {Type: DocumentFieldTypeString, Value: "luke_skywalker"},
	}})
	require.ErrorIs(t, err, ErrPrimaryKeyNotFound)

	err = col.Put(ctx, Document{Fields: map[string]DocumentField{
		"id": {Type: DocumentFieldTypeNumber, Value: 1},
	}})
	require.ErrorIs(t, err, ErrInvalidPrimaryKeyType)

	err = col.Put(ctx, Document{Fields: map[string]DocumentField{
		"id": {Type: DocumentFieldTypeString, Value: 123},
	}})
	require.ErrorIs(t, err, ErrInvalidPrimaryKeyValue)
}

func TestCollection_Delete_NotFound(t *testing.T) {
	ctx := context.Background()
	st := NewStore()

	col, err := st.CreateCollection(ctx, "users", &CollectionConfig{PrimaryKey: "id"})
	require.NoError(t, err)

	err = col.Delete(ctx, "404")
	require.ErrorIs(t, err, ErrDocumentNotFound)
}

func TestCollection_List(t *testing.T) {
	ctx := context.Background()
	st := NewStore()

	col, err := st.CreateCollection(ctx, "users", &CollectionConfig{PrimaryKey: "id"})
	require.NoError(t, err)

	for i := 1; i <= 3; i++ {
		doc := Document{
			Fields: map[string]DocumentField{
				"id": {Type: DocumentFieldTypeString, Value: string(rune('0' + i))},
			},
		}
		require.NoError(t, col.Put(ctx, doc))
	}

	docs := col.List()
	require.Len(t, docs, 3)
}
