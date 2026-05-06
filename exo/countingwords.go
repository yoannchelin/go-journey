package exo

import "strings"


func WordCount(s string) map[string]int {
    result := make(map[string]int)
    s = strings.ToLower(s)
    for _, word := range strings.Fields(s) {
        result[word]++
    }
    return result
}
