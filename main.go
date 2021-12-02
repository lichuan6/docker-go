package main

import "fmt"

func f() {
	fmt.Println("{\"log\": \"123456\"}")
	panic("Hi, this is a panic")
}

func main() {
	fmt.Println("vim-go")
	f()
}
