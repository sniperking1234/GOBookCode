package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filesList := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filesList)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filesList)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			for s, _ := range filesList[line] {
				fmt.Printf("%s\t", s)
			}
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filesList map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		value, ok := filesList[input.Text()]
		if !ok {
			value = make(map[string]int)
			filesList[input.Text()] = value
		}
		value[f.Name()] = 1
	}
	// NOTE: ignoring potential errors from input.Err()
}
