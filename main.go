package main

import (
	"log"
	"my-full-text-search/search"
)

func main() {
	documents, err := search.LoadDocuments("./data.xml")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(len(documents))
}
