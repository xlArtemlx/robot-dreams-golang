package documentstore

type Collection struct {
	cfg       *CollectionConfig
	documents map[string]*Document
}

type CollectionConfig struct {
	PrimaryKey string
}

func (c *Collection) Put(doc Document) error {
	if c.cfg == nil {
		return ErrNilCollectionConfig
	}

	field, ok := doc.Fields[c.cfg.PrimaryKey]
	if !ok {
		return ErrPrimaryKeyNotFound
	}

	if field.Type != DocumentFieldTypeString {
		return ErrInvalidPrimaryKeyType
	}

	key, ok := field.Value.(string)
	if !ok {
		return ErrInvalidPrimaryKeyValue
	}

	c.documents[key] = &doc
	return nil
}

func (c *Collection) Get(key string) (*Document, error) {
	doc, exists := c.documents[key]
	if !exists {
		return nil, ErrDocumentNotFound
	}
	return doc, nil
}

func (c *Collection) Delete(key string) error {
	if _, exists := c.documents[key]; !exists {
		return ErrDocumentNotFound
	}
	delete(c.documents, key)
	return nil
}

func (c *Collection) List() []Document {
	result := make([]Document, 0, len(c.documents))
	for _, doc := range c.documents {
		result = append(result, *doc)
	}
	return result
}
