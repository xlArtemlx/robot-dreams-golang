package documentstore

type Store struct {
	collections map[string]*Collection
}

func NewStore() *Store {
	return &Store{
		collections: make(map[string]*Collection),
	}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (*Collection, error) {
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
	}

	s.collections[name] = col
	return col, nil
}

func (s *Store) GetCollection(name string) (*Collection, error) {
	col, exists := s.collections[name]
	if !exists {
		return nil, ErrCollectionNotFound
	}
	return col, nil
}

func (s *Store) DeleteCollection(name string) error {
	if _, exists := s.collections[name]; !exists {
		return ErrCollectionNotFound
	}
	delete(s.collections, name)
	return nil
}
