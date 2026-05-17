package exo

import "strings"


func IsValidEmail(s string) bool {
	if s == "" {
		return false
	}
	if strings.Contains(s, " ") {
		return false
	}
	atIdx := strings.Index(s, "@")
	if atIdx == 0 || atIdx == len(s)-1 {
		return false
	}
	if strings.Count(s, "@") != 1 {
		return false
	}
	domain := s[atIdx+1:]
	if domain[0] == '.' || !strings.Contains(domain, ".") {
		return false
	}

	return true
}