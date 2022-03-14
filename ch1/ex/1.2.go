//1.3 prints its command and arguments alongside their index
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, " ", arg)
	}
}
