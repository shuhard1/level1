package main

import (
	"fmt"
)

func main() {
	n := 0
	if true {
		n := 1 //объявляет новую переменную, которая затеняет оригинал n во всей области действия оператора if
		n++
	}
	fmt.Println(n) //чтобы здесь вывелось 2 надо в if-е написать n = 1, вместо n := 1
}
