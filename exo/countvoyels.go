package exo

import "strings"

func Voyels(s string) int {
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