package search

import (
	"encoding/xml"
	"errors"
	"os"
	"strings"
)

type Document struct {
	Text string `xml:"abstract"`
	ID   int
}

func Load(path string) (Index, []Document, error) {
	// validate path
	fields := strings.FieldsFunc(path, func(r rune) bool {
		return r == '.'
	})
	if len(fields) < 2 {
		return nil, nil, errors.New("invalid index path")
	}
	indexPath := "./" + fields[0] + ".csv"

	// load documents into memory
	// todo: find a way i don't need to load all the data into memory
	documents, err := loadDocuments("./" + path)
	if err != nil {
		return nil, nil, err
	}

	idx := make(Index)
	// if index file exists, read index from memory instead of data file
	if _, err := os.Stat(indexPath); err == nil {
		err := idx.Load(indexPath)
		return idx, documents, err
	}

	// index does not exist so create and write it
	idx.Add(documents)

	// persist index to disk
	err = idx.Persist(indexPath)
	if err != nil {
		return nil, nil, err
	}

	return idx, documents, err
}

func loadDocuments(path string) ([]Document, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}

	if err := decoder.Decode(&dump); err != nil {
		return nil, err
	}

	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}

	return docs, nil
}
