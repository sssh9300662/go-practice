package main

import "fmt"

func main(){
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// 如果只需要 index，可以簡寫成：
	// for i := range cardValues { ... }
	for i, value := range cardValues {
		fmt.Println(i, value)
	}
}
