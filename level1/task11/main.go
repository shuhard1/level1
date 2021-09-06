package main

import (
	"fmt"

	"github.com/juliangruber/go-intersect"
)

func main() {
	a := []int{1, 2, 3, 7, 8, 9, 0, 6}
	b := []int{2, 3, 4, 6, 8}
	fmt.Println(intersect.Simple(a, b))
}
