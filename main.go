package main

import (
    "fmt"
    "lesson_3/documentstore"
)

func main() {

	doc := &documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "doc_1",
			},
		},
	}

	documentstore.Put(doc)

	retrievedDoc, found := documentstore.Get("doc_1")
	if found {
		fmt.Println("Document found:", retrievedDoc)
	} else {
		fmt.Println("Document not found")
	}

	documentstore.Delete("doc_1")
}