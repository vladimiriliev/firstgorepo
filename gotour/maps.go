package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s);
	result := make (map[string]int)
	for _, value := range words {
		elem, ok := result[value]
		if ok {
			result[value]=elem+1
		} else {
			result[value]=1
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
