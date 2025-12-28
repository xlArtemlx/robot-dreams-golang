package documentstore

import "errors"

var (
	///store/collection/document errors
	ErrDocumentNotFound         = errors.New("document not found")
	ErrCollectionAlreadyExists  = errors.New("collection already exists")
	ErrCollectionNotFound       = errors.New("collection not found")
	ErrUnsupportedDocumentField = errors.New("unsupported document field")

	//validation errors
	ErrNilCollectionConfig    = errors.New("collection config is null")
	ErrEmptyCollectionName    = errors.New("collection name is empty")
	ErrPrimaryKeyNotFound     = errors.New("primary key field not found")
	ErrInvalidPrimaryKeyType  = errors.New("primary key must be a string field")
	ErrInvalidPrimaryKeyValue = errors.New("primary key value must be a strings")

	// Marshal errors
	ErrNilDocument = errors.New("document is nil")
	ErrNilOutput   = errors.New("output is nil")
)
