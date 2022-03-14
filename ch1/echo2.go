//echo2 prints its command-line arguments using a different for loop
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg

		sep = " "
	}
	fmt.Println(s)
}
