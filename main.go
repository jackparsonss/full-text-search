package main

import (
	"log"

	"github.com/jackparsonss/full-text-search/search"
	"github.com/jackparsonss/full-text-search/web"
)

func main() {
	log.Println("Loading documents...")
	index, documents, err := search.Load("data.xml")
	if err != nil {
		log.Fatal(err)
	}
	web.Entry(index, documents)
}
