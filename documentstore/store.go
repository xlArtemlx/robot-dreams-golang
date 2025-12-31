package documentstore

import (
	"context"
	"log/slog"
)

type Store struct {
	collections map[string]*Collection
	log         *slog.Logger
}

func NewStore() *Store {
	return &Store{
		collections: make(map[string]*Collection),
		log:         Logger(),
	}
}

func (s *Store) CreateCollection(ctx context.Context, name string, cfg *CollectionConfig) (*Collection, error) {
	if name == "" {
		return nil, ErrEmptyCollectionName
	}
	if cfg == nil {
		return nil, ErrNilCollectionConfig
	}

	if _, exists := s.collections[name]; exists {
		return nil, ErrCollectionAlreadyExists
	}

	col := &Collection{
		cfg:       cfg,
		documents: make(map[string]*Document),
		log:       s.log.With("collection", name),
	}

	s.collections[name] = col

	s.log.InfoContext(ctx, "collection created",
		"event", "collection.create",
		"collection", name,
		"primary_key", cfg.PrimaryKey,
	)
	return col, nil
}

func (s *Store) GetCollection(name string) (*Collection, error) {
	col, exists := s.collections[name]
	if !exists {
		return nil, ErrCollectionNotFound
	}
	return col, nil
}

func (s *Store) DeleteCollection(ctx context.Context, name string) error {
	if _, exists := s.collections[name]; !exists {
		return ErrCollectionNotFound
	}
	delete(s.collections, name)
	s.log.InfoContext(ctx, "collection deleted",
		"event", "collection.delete",
		"collection", name,
	)
	return nil
}
