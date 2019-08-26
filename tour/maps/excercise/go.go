package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	retval := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		count, ok := retval[word]
		if ok {
			retval[word] = count + 1
		} else {
			retval[word] = 1
		}
	}

	return retval
}

func main() {
	wc.Test(WordCount)
}
