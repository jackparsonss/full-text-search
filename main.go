package main

import (
	"log"

	"github.com/jackparsonss/full-text-search/search"
)

func main() {
	documents, err := search.LoadDocuments("./data.xml")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(len(documents))
}
