package backend

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findFirstStringInBracket(t *testing.T) {
	input := "||abcdef{(luqman )  abcdef"
	expected := "luqman"
	assert.Equal(t, expected, findFirstStringInBracket(input))

}
