package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Println(fToc(freezingF))
	fmt.Println(boilingF)
}

func fToc(f float64) float64 {
	return (f - 32)
}
