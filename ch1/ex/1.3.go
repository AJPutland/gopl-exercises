package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	t1 := time.Now()
	//First approach (inefficient string usage)
	s := ""
	for _, x := range os.Args[1:] {
		s += x + " "
	}
	fmt.Println(s, time.Now().Sub(t1))

	t2 := time.Now()
	//Second approach (efficient string usage)
	fmt.Println(strings.Join(os.Args[1:], " "), time.Now().Sub(t2))

	t3 := time.Now()
	//Third approach using string builder (efficient string usage)
	var b strings.Builder
	for _, y := range os.Args[1:] {
		b.WriteString(y + " ")
	}
	fmt.Println(b.String(), time.Now().Sub(t3))

}
