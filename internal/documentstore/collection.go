package documentstore

import (
	"context"
	"log/slog"
)

type Collection struct {
	name      string
	cfg       *CollectionConfig
	documents map[string]*Document
	log       *slog.Logger
}

type CollectionConfig struct {
	PrimaryKey string
}

func (c *Collection) Put(ctx context.Context, doc Document) error {
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

	_, existed := c.documents[key]
	c.documents[key] = &doc

	event := "document.create"
	msg := "document created"
	if existed {
		event = "document.update"
		msg = "document updated"
	}

	logger := c.log
	if logger == nil {
		logger = slog.Default()
	}

	logger.InfoContext(ctx, msg,
		"event", event,
		"collection", c.name,
		"doc_id", key,
		"primary_key", c.cfg.PrimaryKey,
	)

	return nil
}

func (c *Collection) Get(key string) (*Document, error) {
	doc, exists := c.documents[key]
	if !exists {
		return nil, ErrDocumentNotFound
	}
	return doc, nil
}

func (c *Collection) Delete(ctx context.Context, key string) error {
	if _, exists := c.documents[key]; !exists {
		return ErrDocumentNotFound
	}
	delete(c.documents, key)

	logger := c.log
	if logger == nil {
		logger = slog.Default()
	}

	logger.InfoContext(ctx, "document deleted",
		"event", "document.delete",
		"collection", c.name,
		"doc_id", key,
	)
	return nil
}

func (c *Collection) List() []Document {
	result := make([]Document, 0, len(c.documents))
	for _, doc := range c.documents {
		result = append(result, *doc)
	}
	return result
}
