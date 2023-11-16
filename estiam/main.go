package main

import (
	"bufio"
	"estiam/dictionary"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	dict := dictionary.New("dictionary.json")

	for {
		fmt.Println("Enter command (add, define, remove, list, exit):")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "add":
			actionAdd(dict, reader)
		case "define":
			actionDefine(dict, reader)
		case "remove":
			actionRemove(dict, reader)
		case "list":
			actionList(dict)
		case "exit":
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}

func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter word: ")
	word, _ := reader.ReadString('\n')
	fmt.Print("Enter definition: ")
	definition, _ := reader.ReadString('\n')

	d.Add(strings.TrimSpace(word), strings.TrimSpace(definition))
	fmt.Println("Added.")
}

func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter word: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	entry, err := d.Get(word)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Definition:", entry)
}

func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter word: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	d.Remove(word)
	fmt.Println("Removed.")
}

func actionList(d *dictionary.Dictionary) {
	words, _ := d.List()
	for _, word := range words {
		fmt.Println(word)
	}
}
