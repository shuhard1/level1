package main

import (
	"fmt"
)

func reverse(s string) string { //rune() -это не функция, это синтаксис для преобразования типа в rune
	chars := []rune(s) ////преобразование строк в срезы
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i] //элементы меняются местами по передниму и задниму индексу
	}
	return string(chars) //возваращает строку
}

func main() {
	fmt.Println(reverse("snow dog sun - sun dog snow")) //запуск функции
}
