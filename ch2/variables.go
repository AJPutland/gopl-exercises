//Variables explores all of the different ways you can declare / initialize variables in Go
package main

import (
	"fmt"
	"os"
)

//Package level variables
var a int = 1
var b = 2 // type inferred
var c int // zero value used

func main() {
	fmt.Printf("Package level variables\na:%d\nb:%d\nc:%d\n", a, b, c)

	//Local variables and multiple declarations / intializations
	var d = 3
	var e, f, g int             //multiple variables of same type
	var h, i, j = 4, "hi", true //multiple variables with assgnment (no type included to allow different types)

	fmt.Printf("Function level variables\nd:%d\ne:%d\nf:%d\ng:%d\nh:%d\ni:%s\nj:%v\n", d, e, f, g, h, i, j)

	//Initializing a set of variables using a function that returns multiuple values
	var k, l = os.Open("non-existent")
	fmt.Printf("Using a function to intialize a set of variables\nk:%v\nl:%v\n", k, l)

	//Using short variable declarations
	m := 9
	n, o := "bye", 5
	p, q := os.Open("non-existent")
	fmt.Printf("Using short variable declaration\nm:%v\nn:%v\no:%v\np:%v\nq:%v\n", m, n, o, p, q)

	//Short variable declaration can act like an assignment, but needs to declare at least one
	r, s := 6, 7
	t, s := 8, 9
	fmt.Printf("Short variable declarations can re assign\nr:%v\ns:%v\nt:%v\n", r, s, t)
}
