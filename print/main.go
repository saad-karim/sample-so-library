package main

import (
	"C"
)

import (
	"fmt"
)

//export PrintHello
func PrintHello() {
	fmt.Println("Hello!")
}

func main() {}
