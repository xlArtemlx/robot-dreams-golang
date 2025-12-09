package documentstore

type Collection struct {
	cfg       *CollectionConfig
	documents map[string]*Document
}

type CollectionConfig struct {
	PrimaryKey string
}

func (s *Collection) Put(doc Document) {
	field, ok := doc.Fields[s.cfg.PrimaryKey]
	if !ok {
		return
	}

	if field.Type != DocumentFieldTypeString {
		return
	}

	key, ok := field.Value.(string)
	if !ok {
		return
	}
	s.documents[key] = &doc
}

func (s *Collection) Get(key string) (*Document, bool) {
	doc, exists := s.documents[key]
	return doc, exists
}

func (s *Collection) Delete(key string) bool {
	if _, exists := s.documents[key]; exists {
		delete(s.documents, key)
		return true
	}
	return false
}

func (s *Collection) List() []Document {
	result := make([]Document, 0, len(s.documents))
	for _, doc := range s.documents {
		result = append(result, *doc)
	}
	return result
}
