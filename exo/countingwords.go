package exo

import "strings"

func WordCount(s string) map[string]int {
    res := make(map[string]int)
    s = strings.ToLower(s)

    for _, word := range strings.Fields(s) {
        res[word]++
    }
    return res
} 