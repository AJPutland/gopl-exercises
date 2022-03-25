// Dup3 reads in files and prints the number of duplicated lines
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, arg := range os.Args[1:] {
		file, err := os.ReadFile(arg)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Dup3 %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(file), "\n") {
			counts[line]++
		}
	}

	for k, v := range counts {
		if v > 1 {
			fmt.Println(v, k)
		}
	}
}
