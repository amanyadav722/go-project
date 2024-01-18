package dictionary

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	dict := New("test_dictionary.json")
	defer os.Remove("test_dictionary.json")

	err := dict.Add("test", "a test word")
	assert.NoError(t, err)

	got, err := dict.Get("test")
	assert.NoError(t, err)
	assert.Equal(t, "a test word", got)
}
