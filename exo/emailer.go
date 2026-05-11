package exo

import (
	"strings"
)

func IsValidEmail(s string) bool {
	if s == "" {
		return false
	}
	if strings.Count(s, "@") != 1 {
		return false
	}
	adx := strings.Index(s, "@")
	if adx == 0 || adx == len(s)-1 {
		return false
	}
	domain := s[adx+1:]
    if !strings.Contains(domain, ".") {
        return false
    }
	if strings.Contains(s, " ") {
		return false
	}

	return true
}