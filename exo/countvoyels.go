package exo

import "strings"

func CountVowels(s string) int {
	voyels := "aeiouy"
	s = strings.ToLower(s)
	count := 0
	for _, r := range s {
		if strings.ContainsRune(voyels, r) {
			count++
		}
	}
	return count
}