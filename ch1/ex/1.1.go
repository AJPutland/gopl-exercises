// 1.1 prints its command-line arguments and name of command out using the string.Join function
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
