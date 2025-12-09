package documentstore

type Store struct {
	collections map[string]*Collection
}

func NewStore() *Store {
	return &Store{
		collections: make(map[string]*Collection),
	}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (bool, *Collection) {
	if _, exists := s.collections[name]; exists {
		return false, nil
	}

	col := &Collection{
		cfg:       cfg,
		documents: make(map[string]*Document),
	}

	s.collections[name] = col

	return true, col
}

func (s *Store) GetCollection(name string) (*Collection, bool) {
	col, exists := s.collections[name]
	return col, exists
}

func (s *Store) DeleteCollection(name string) bool {
	if _, exists := s.collections[name]; exists {
		delete(s.collections, name)
		return true
	}
	return false
}
