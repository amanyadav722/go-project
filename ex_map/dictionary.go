package main

import "sort"

type Dictionary struct {
	entries map[string]string
}

func NewDictionary() *Dictionary {
	return &Dictionary{entries: make(map[string]string)}
}

func (d *Dictionary) Add(word, definition string) {
	d.entries[word] = definition
}

func (d *Dictionary) Get(word string) (string, bool) {
	definition, exists := d.entries[word]
	return definition, exists
}

func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
}

func (d *Dictionary) List() []string {
	var words []string
	for word := range d.entries {
		words = append(words, word)
	}
	sort.Strings(words)
	return words
}
