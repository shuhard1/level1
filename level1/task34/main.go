package main

import (
	"fmt"
)

func main() {
	s := "abcdefghjklmnpq"

	checkUnique(&s)
}

func checkUnique(s *string) bool { //функция, которая принимает ссылку *string

	var a [256]bool            //создаем массив с длинной 256
	for _, ascii := range *s { //проходимся по массиву s

		if a[ascii] { //если a[ascii] равен true, то
			fmt.Println("в строке есть повторяющиеся символы")
			return false //возврат false
		}
		a[ascii] = true //a[ascii] присваевается значение true
	}
	fmt.Println("все символы в строке уникальные")
	return true
}
