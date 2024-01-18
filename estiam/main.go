package main

import (
	"bufio"
	"encoding/json"
	"estiam/dictionary"
	"estiam/middleware"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	dict := dictionary.New("dictionary.json")

	router := mux.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.AuthMiddleware)

	router.HandleFunc("/word", addWord(dict)).Methods("POST")
	router.HandleFunc("/word/{word}", getDefinition(dict)).Methods("GET")
	router.HandleFunc("/words", getAllWords(dict)).Methods("GET")
	router.HandleFunc("/word/{word}", deleteWord(dict)).Methods("DELETE")

	go func() {
		fmt.Println("Server is up and running here on http://localhost:8081")
		err := http.ListenAndServe(":8081", router)
		if err != nil {
			log.Fatalf("Failed to start server just go and fix it first: %v", err)
		}
	}()

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
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}

		if len(payload.Word) < 1 || len(payload.Word) > 50 {
			http.Error(w, "The word must be between 1 and 50 characters long", http.StatusBadRequest)
			return
		}
		if len(payload.Definition) < 1 || len(payload.Definition) > 500 {
			http.Error(w, "The definition must be between 1 and 200 characters long", http.StatusBadRequest)
			return
		}

		if err := d.Add(payload.Word, payload.Definition); err != nil {
			http.Error(w, "Failed to add the word to the dictionary", http.StatusInternalServerError)
			log.Printf("Add: %v", err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Your word added successfully")
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

func getAllWords(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		entries, err := d.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(entries); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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
		json.NewEncoder(w).Encode("Your stuff deleted successfully")
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
