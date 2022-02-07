package main

import "fmt"

func main() {
	hello := "Hello World"
	helloSlice := string2ByteSlice(hello)
	fmt.Println(helloSlice)

	helloString := byteSlice2String(helloSlice)
	fmt.Println(helloString)
}

// byte slice 轉成 string
func byteSlice2String(b []byte) string {
	return string(b)
}

// string 轉成 byte slice
func string2ByteSlice(s string) []byte {
	return []byte(s)
}
