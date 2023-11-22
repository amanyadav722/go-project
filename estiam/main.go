package main

import (
	"bufio"
	"encoding/json"
	"estiam/dictionary"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	dict := dictionary.New("dictionary.json")

	router := mux.NewRouter()

	fmt.Println("Server is running on http://localhost:8080")

	router.HandleFunc("/word", addWord(dict)).Methods("POST")
	router.HandleFunc("/word/{word}", getDefinition(dict)).Methods("GET")
	router.HandleFunc("/word/{word}", deleteWord(dict)).Methods("DELETE")

	http.ListenAndServe(":8080", router)

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

func addWord(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Word       string `json:"word"`
			Definition string `json:"definition"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := d.Add(payload.Word, payload.Definition); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Entry added successfully")
	}
}

func getDefinition(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		word := vars["word"]

		definition, err := d.Get(word)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(definition)
	}
}

func deleteWord(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		word := vars["word"]

		if err := d.Remove(word); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Entry deleted successfully")
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
