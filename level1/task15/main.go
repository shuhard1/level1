package main

import (
	"fmt"
)

func main() {
	a := "second"
	b := "first"
	fmt.Println(a, b)
	b, a = a, b
	fmt.Println(a, b)
}
