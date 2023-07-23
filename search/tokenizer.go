package search

import (
	"strings"
	"unicode"

	snowballeng "github.com/kljensen/snowball/english"
)

func Analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowerCaseFilter(tokens)
	tokens = stopWordsFilter(tokens)
	tokens = stemmerFilter(tokens)

	return tokens
}

// split tokens
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

// filters
func lowerCaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}

	return r
}

// top ten stop words to remove
var stopWords = map[string]struct{}{
	"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
	"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
}

func stopWordsFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for _, token := range tokens {
		if _, ok := stopWords[token]; !ok {
			r = append(r, token)
		}
	}

	return r
}

// stem words
func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowballeng.Stem(token, false)
	}

	return r
}
