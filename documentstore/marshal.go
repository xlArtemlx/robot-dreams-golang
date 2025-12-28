package documentstore

import (
	"encoding/json"
)

func MarshalDocument(input any) (*Document, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	fields := make(map[string]DocumentField, len(m))
	for k, v := range m {
		field, err := buildField(v)
		if err != nil {
			return nil, err
		}
		fields[k] = field
	}

	return &Document{Fields: fields}, nil
}

func UnmarshalDocument(doc *Document, output any) error {
	if doc == nil {
		return ErrNilDocument
	}
	if output == nil {
		return ErrNilOutput
	}

	m := make(map[string]any, len(doc.Fields))
	for k, f := range doc.Fields {
		v, err := fieldToAny(f)
		if err != nil {
			return err
		}
		m[k] = v
	}

	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, output)
}

func buildField(v any) (DocumentField, error) {
	switch t := v.(type) {
	case string:
		return DocumentField{Type: DocumentFieldTypeString, Value: t}, nil
	case bool:
		return DocumentField{Type: DocumentFieldTypeBool, Value: t}, nil
	case float64:
		return DocumentField{Type: DocumentFieldTypeNumber, Value: t}, nil
	case []any:
		items := make([]DocumentField, 0, len(t))
		for _, item := range t {
			f, err := buildField(item)
			if err != nil {
				return DocumentField{}, err
			}
			items = append(items, f)
		}
		return DocumentField{Type: DocumentFieldTypeArray, Value: items}, nil
	case map[string]any:
		obj := make(map[string]DocumentField, len(t))
		for k, val := range t {
			f, err := buildField(val)
			if err != nil {
				return DocumentField{}, err
			}
			obj[k] = f
		}
		return DocumentField{Type: DocumentFieldTypeObject, Value: obj}, nil
	case nil:
		return DocumentField{}, ErrUnsupportedDocumentField
	default:
		return DocumentField{}, ErrUnsupportedDocumentField
	}
}

func fieldToAny(f DocumentField) (any, error) {
	switch f.Type {
	case DocumentFieldTypeString:
		v, ok := f.Value.(string)
		if !ok {
			return nil, ErrUnsupportedDocumentField
		}
		return v, nil
	case DocumentFieldTypeBool:
		v, ok := f.Value.(bool)
		if !ok {
			return nil, ErrUnsupportedDocumentField
		}
		return v, nil
	case DocumentFieldTypeNumber:
		switch v := f.Value.(type) {
		case float64:
			return v, nil
		case float32:
			return float64(v), nil
		case int:
			return float64(v), nil
		case int64:
			return float64(v), nil
		default:
			return nil, ErrUnsupportedDocumentField
		}
	case DocumentFieldTypeArray:
		items, ok := f.Value.([]DocumentField)
		if !ok {
			return nil, ErrUnsupportedDocumentField
		}
		out := make([]any, 0, len(items))
		for _, it := range items {
			v, err := fieldToAny(it)
			if err != nil {
				return nil, err
			}
			out = append(out, v)
		}
		return out, nil
	case DocumentFieldTypeObject:
		obj, ok := f.Value.(map[string]DocumentField)
		if !ok {
			return nil, ErrUnsupportedDocumentField
		}
		out := make(map[string]any, len(obj))
		for k, it := range obj {
			v, err := fieldToAny(it)
			if err != nil {
				return nil, err
			}
			out[k] = v
		}
		return out, nil
	default:
		return nil, ErrUnsupportedDocumentField
	}
}
