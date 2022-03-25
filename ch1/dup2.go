// Dup2 reads from either stdin or files and prints out any duplicate lines
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		//count lines from stdin
		countLines(os.Stdin, counts)

	} else {
		//count line using files
		for _, x := range files {
			file, err := os.Open(x)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %s\n", err)
				continue
			}

			countLines(file, counts)
			file.Close()
		}
	}

	// print out duplicate lines
	for k, v := range counts {
		if v > 1 {
			fmt.Println(v, ":", k)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
