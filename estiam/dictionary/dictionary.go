package dictionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Dictionary struct {
	file string
}

func New(file string) *Dictionary {
	return &Dictionary{file: file}
}

func (d *Dictionary) load() (map[string]string, error) {
	var data map[string]string
	fileData, err := ioutil.ReadFile(d.file)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]string), nil
		}
		return nil, err
	}
	err = json.Unmarshal(fileData, &data)
	return data, err
}

func (d *Dictionary) save(data map[string]string) error {
	fileData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(d.file, fileData, 0644)
}

func (d *Dictionary) Add(word, definition string) error {
	data, err := d.load()
	if err != nil {
		return err
	}
	data[word] = definition
	return d.save(data)
}

func (d *Dictionary) Get(word string) (string, error) {
	data, err := d.load()
	if err != nil {
		return "", err
	}
	definition, exists := data[word]
	if !exists {
		return "", fmt.Errorf("word not found: %s", word)
	}
	return definition, nil
}

func (d *Dictionary) Remove(word string) error {
	data, err := d.load()
	if err != nil {
		return err
	}
	delete(data, word)
	return d.save(data)
}

func (d *Dictionary) List() ([]string, error) {
	data, err := d.load()
	if err != nil {
		return nil, err
	}
	var words []string
	for word := range data {
		words = append(words, word)
	}
	return words, nil
}
