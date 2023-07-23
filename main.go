package main

import (
	"log"

	"github.com/jackparsonss/full-text-search/search"
)

func main() {
	log.Println("Loading documents...")
	documents, err := search.LoadDocuments("./data.xml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Loaded: ", len(documents), " documents")

	idx := make(search.Index)

	log.Println("Creating index...")
	idx.Add(documents)
	log.Println("Created index")

	log.Println("Searching...")
	res := idx.Search("Small wild cat")
	log.Println("Search Result Index: ", res)
	log.Println("Search Result Text: ", documents[res[0]].Text)
}
