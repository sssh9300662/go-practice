package main

import "fmt"

func main() {
	// 沒有參數
	func() {
		fmt.Println("Hello anonymous")
	}()

	// 有參數
	func(i, j int) {
		fmt.Println(i + j)
	}(1, 2)
}
