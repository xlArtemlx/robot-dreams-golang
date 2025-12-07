package documentstore

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType  = "string"
 	DocumentFieldTypeNumber DocumentFieldType  = "number"
 	DocumentFieldTypeBool   DocumentFieldType  = "bool"
 	DocumentFieldTypeArray  DocumentFieldType  = "array"
 	DocumentFieldTypeObject DocumentFieldType  = "object"
)

type DocumentField struct {
	Type  DocumentFieldType
	Value interface{}
}

type Document struct {
 	Fields map[string]DocumentField
}

var documents = map[string]*Document{}

func Put(doc *Document) {
	field, emptyKey := doc.Fields["key"]
	if !emptyKey {
		return
	}

	if field.Type != DocumentFieldTypeString {
		return
	}

	key, ok := field.Value.(string)
	if !ok {
		return
	}
	documents[key] = doc
}

func Get(key string) (*Document, bool) {
	return documents[key], documents[key] != nil
}

func Delete(key string) bool { 
	if _, exists := documents[key]; exists {
		delete(documents, key)
		return true
	}
	return false
}

func List() []*Document {
	result := make([]*Document, 0, len(documents))
	for _, doc := range documents {
		result = append(result, doc)
	}
	return result
}
