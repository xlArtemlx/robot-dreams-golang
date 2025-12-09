package main

import (
	"fmt"
	"lesson_4/documentstore"
)

func main() {

	store := documentstore.NewStore()

	created, collection := store.CreateCollection("users", &documentstore.CollectionConfig{
		PrimaryKey: "key",
	})

	if !created {
		fmt.Println("Collection already exists")
		return
	}

	doc := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "doc_1",
			},
			"name": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "John",
			},
		},
	}

	collection.Put(doc)

	retrieved, found := collection.Get("doc_1")
	if found {
		fmt.Println("Document found:", retrieved)
	} else {
		fmt.Println("Document not found")
	}

	deleted := collection.Delete("doc_1")
	fmt.Println("Document deleted:", deleted)

	fmt.Println("List:", collection.List())
}
