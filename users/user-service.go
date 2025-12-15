package users

import (
	"errors"

	"lesson_5/documentstore"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Service struct {
	coll *documentstore.Collection
}

func UserService(st *documentstore.Store) (*Service, error) {
	coll, err := st.CreateCollection("users", &documentstore.CollectionConfig{PrimaryKey: "id"})
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
	return &Service{coll: coll}, nil
}

func (s *Service) CreateUser(id string, name string) (*User, error) {
	user := &User{ID: id, Name: name}

	doc, err := documentstore.MarshalDocument(user)
	if err != nil {
		return nil, err
	}

	if err := s.coll.Put(*doc); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) ListUsers() ([]User, error) {
	docs := s.coll.List()
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

func (s *Service) GetUser(userID string) (*User, error) {
	doc, err := s.coll.Get(userID)
	if err != nil {
		if errors.Is(err, documentstore.ErrDocumentNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	user := &User{}
	if err := documentstore.UnmarshalDocument(doc, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) DeleteUser(userID string) error {
	if err := s.coll.Delete(userID); err != nil {
		if errors.Is(err, documentstore.ErrDocumentNotFound) {
			return ErrUserNotFound
		}
		return err
	}
	return nil
}
