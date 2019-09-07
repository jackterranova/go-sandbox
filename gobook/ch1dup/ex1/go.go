package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for filename, dupcount := range counts {
		for line, n := range dupcount {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", filename, n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) { //maps params are references, not a copy of the whole map
	input := bufio.NewScanner(f)

	counts[f.Name()] = make(map[string]int)

	for input.Scan() {
		counts[f.Name()][input.Text()]++
	}
}
