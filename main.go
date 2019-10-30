package main

import (
	"fmt"
	"log"

	"gopkg.in/jdkato/prose.v2"
)

func main() {
	// Create a new document with the default configuration:
	sentence := "This was taken during the Easter celebrations at Real de Catorce, MX."
	doc, err := prose.NewDocument(sentence)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the doc's tokens:
	fmt.Println(sentence)
	for _, tok := range doc.Tokens() {
		fmt.Println(tok.Text, tok.Tag, tok.Label)
		// Go NNP B-GPE
		// is VBZ O
		// an DT O
		// ...
	}

	fmt.Println("Entities")
	// Iterate over the doc's named-entities:
	for _, ent := range doc.Entities() {
		fmt.Println(ent.Text, ent.Label)
		// Go GPE
		// Google GPE
	}
	fmt.Println("Sentences")
	// Iterate over the doc's sentences:
	for _, sent := range doc.Sentences() {
		fmt.Println(sent.Text)
		// Go is an open-source programming language created at Google.
	}
	// prose.ModelFromData(name)
	// prose.LabeledEntity
}
