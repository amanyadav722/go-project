package main

import (
	"bytes"
	"estiam/dictionary"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddWordHandler(t *testing.T) {
	dict := dictionary.New("test_dictionary.json")
	defer os.Remove("test_dictionary.json")

	router := mux.NewRouter()
	router.HandleFunc("/word", addWord(dict)).Methods("POST")

	payload := []byte(`{"word":"test","definition":"a test word"}`)
	req, _ := http.NewRequest("POST", "/word", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestGetDefinitionHandler(t *testing.T) {
	dict := dictionary.New("test_dictionary.json")
	defer os.Remove("test_dictionary.json")

	router := mux.NewRouter()
	router.HandleFunc("/word/{word}", getDefinition(dict)).Methods("GET")

	dict.Add("hello", "a greeting")

	req, _ := http.NewRequest("GET", "/word/hello", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestDeleteWordHandler(t *testing.T) {
	dict := dictionary.New("test_dictionary.json")
	defer os.Remove("test_dictionary.json")

	err := dict.Add("testword", "a test definition")
	require.NoError(t, err)

	router := mux.NewRouter()
	router.HandleFunc("/word/{word}", deleteWord(dict)).Methods("DELETE")

	req, _ := http.NewRequest("DELETE", "/word/testword", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	_, err = dict.Get("testword")
	assert.Error(t, err)
}

func TestGetNonExistentWord(t *testing.T) {
	dict := dictionary.New("test_dictionary.json")
	defer os.Remove("test_dictionary.json")

	router := mux.NewRouter()
	router.HandleFunc("/word/{word}", getDefinition(dict)).Methods("GET")

	req, _ := http.NewRequest("GET", "/word/nonexistentword", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}
