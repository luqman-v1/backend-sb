package backend

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func Anagram(input []string) [][]string {
	m := make(map[string][]string, 0)
	for i := 0; i < len(input); i++ {
		v, ok := m[SortString(input[i])]
		if ok {
			m[SortString(input[i])] = append(v, input[i])
		} else {
			m[SortString(input[i])] = append([]string{}, input[i])
		}
	}
	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
