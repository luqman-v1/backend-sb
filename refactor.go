package backend

import "strings"

type Rune struct {
	str []rune
}

func New() *Rune {
	return &Rune{}
}

func subStrIndex(input, substr string) int {
	return strings.Index(input, substr)
}

//subRune
func (r *Rune) subRune(input string, indexOpen, indexClose int) *Rune {
	runes := []rune(input)
	r.str = runes[indexOpen:indexClose]
	return r
}

func (r *Rune) toString() string {
	return string(r.str)
}

func findFirstStringInBracket(str string) string {
	n := New()
	bracketOpen := "("
	bracketClose := ")"
	totalStr := len(str)
	if totalStr > 0 {
		indexFirstBracketFound := subStrIndex(str, bracketOpen)
		if indexFirstBracketFound >= 0 {
			wordsAfterFirstBracket := n.subRune(str, indexFirstBracketFound, totalStr).toString()
			indexClosingBracketFound := subStrIndex(wordsAfterFirstBracket, bracketClose)
			if indexClosingBracketFound >= 0 {
				return n.subRune(wordsAfterFirstBracket, 1, indexClosingBracketFound-1).toString()
			}
		}
	}
	return ""
}
