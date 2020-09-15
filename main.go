package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var inputFile string
	var patternFile string

	var patterns []string

	var out *os.File
	var err error

	flag.StringVar(&inputFile, "input", "", "Input file (stdin if not specified)")
	flag.StringVar(&patternFile, "pattern", "", "Pattern file")

	flag.Parse()

	if patternFile != "" {
		pf, err := os.Open(patternFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}

		sc := bufio.NewScanner(pf)

		for sc.Scan() {
			pattern := sc.Text()
			if pattern != "" {
				patterns = append(patterns, pattern)
			}
		}
	} else {
		flag.Usage()
		return
	}

	if inputFile != "" {
		out, err = os.Open(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
	} else {
		out = os.Stdin
	}

	in := bufio.NewScanner(out)

	for in.Scan() {
		line := in.Text()

		patternFound := false

		for _, pattern := range patterns {
			if strings.Contains(line, pattern) {
				patternFound = true
				break
			}
		}

		if !patternFound {
			fmt.Println(line)
		}
	}
}
