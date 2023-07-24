package main

import (
	"log"

	"github.com/jackparsonss/full-text-search/search"
)

func main() {
	log.Println("Loading documents...")
	index, _, err := search.Load("data.xml")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Searching...")
	res := index.Search("Small wild cat")
	log.Println("Search Result Index: ", res)
}
