package main

import "fmt"

func main() {
	dict := NewDictionary()

	dict.Add("Go", "A statically typed compiled programming language designed at Google.")
	dict.Add("Map", "A collection type that holds key-value pairs")

	word := "Go"
	definition, exists := dict.Get(word)
	if exists {
		fmt.Printf("Definition of '%s': %s\n", word, definition)
	} else {
		fmt.Println("Word not found:", word)
	}

	dict.Remove("Map")

	fmt.Println("Words in dictionary:")
	for _, w := range dict.List() {
		def, _ := dict.Get(w)
		fmt.Printf("%s: %s\n", w, def)
	}
}
