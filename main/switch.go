package main

import "fmt"

func main() {
	num := 1
	switch start := 0; num {
	case start:
		fmt.Println("0")
	case start + 1:
		fmt.Println("1")
	default:
		fmt.Println("other")
	}
}
