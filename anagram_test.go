package backend

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	input := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	expected := [][]string{
		{"kita", " atik", "tika"},
		{"aku", "kua"},
		{"makan"},
		{"kia"},
	}
	assert.Equal(t, expected, expected, Anagram(input))
}
