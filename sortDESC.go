package sortDESC

import (
	"errors"
	"sort"
	"strconv"
)

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

func Sort(wording string) (string, error) {

	if _, err := strconv.ParseFloat(wording, 64); err == nil {
		return "", errors.New("error, your input look like a number")

	}
	newWording := SortString(wording)
	return newWording, nil

}
