package search

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Index map[string][]int

func (idx Index) Search(text string) []int {
	var r []int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
				continue
			}
			r = intersection(r, ids)
		} else {
			return nil // no token
		}
	}

	return r
}

func (idx Index) Add(docs []Document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue // skip duplicates
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

func (idx Index) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, row := range data {
		fields := strings.Fields(row[1])
		str := row[0]
		for _, v := range fields {
			val, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			idx[str] = append(idx[str], val)
		}
	}

	return nil
}

func (idx Index) Persist(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var data [][]string
	for key, value := range idx {
		row := []string{key, strings.Trim(strings.Replace(fmt.Sprint(value), " ", " ", -1), "[]")}
		data = append(data, row)
	}
	writer.WriteAll(data)

	return nil
}

func intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)

	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}

	return r
}
