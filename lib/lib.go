package lib

import "github.com/mdwhatcott/go-set/v2/set"

func IsMatch(letters, word string) bool {
	l := set.Of([]rune(letters)...)
	for _, c := range word {
		if !l.Contains(c) {
			return false
		}
		l.Remove(c)
	}
	return true
}
func FindMatches(letters string, dict ...string) (results []string) {
	for _, word := range dict {
		if 2 < len(word) && len(word) <= len(letters) {
			if IsMatch(letters, word) {
				results = append(results, word)
			}
		}
	}
	return results
}
