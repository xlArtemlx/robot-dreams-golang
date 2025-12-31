package users

import (
	"context"
	"errors"

	"github.com/xlArtemlx/robot-dreams-golang/documentstore"
)

type DocumentStoreUserRepository struct {
	coll *documentstore.Collection
}

func NewDocumentStoreUserRepository(ctx context.Context, st *documentstore.Store) (*DocumentStoreUserRepository, error) {
	coll, err := st.CreateCollection(ctx, "users", &documentstore.CollectionConfig{PrimaryKey: "id"})
	if err != nil {
		if errors.Is(err, documentstore.ErrCollectionAlreadyExists) {
			coll, err = st.GetCollection("users")
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &DocumentStoreUserRepository{coll: coll}, nil
}

func (r *DocumentStoreUserRepository) Create(ctx context.Context, user *User) error {
	doc, err := documentstore.MarshalDocument(user)
	if err != nil {
		return err
	}

	return r.coll.Put(ctx, *doc)
}

func (r *DocumentStoreUserRepository) List() ([]User, error) {

	docs := r.coll.List()
	result := make([]User, 0, len(docs))

	for _, d := range docs {
		var user User
		if err := documentstore.UnmarshalDocument(&d, &user); err != nil {
			return nil, err
		}
		result = append(result, user)
	}

	return result, nil
}

func (r *DocumentStoreUserRepository) GetByID(id string) (*User, error) {

	doc, err := r.coll.Get(id)
	if err != nil {
		if errors.Is(err, documentstore.ErrDocumentNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	var user User
	if err := documentstore.UnmarshalDocument(doc, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *DocumentStoreUserRepository) Delete(ctx context.Context, id string) error {
	if err := r.coll.Delete(ctx, id); err != nil {
		if errors.Is(err, documentstore.ErrDocumentNotFound) {
			return ErrUserNotFound
		}
		return err
	}
	return nil
}
