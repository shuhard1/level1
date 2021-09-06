package main

import (
	"fmt"
)

func Reverse(s string) string { //rune() -это не функция, это синтаксис для преобразования типа в rune
	runes := []rune(s) //преобразование строк в срезы
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] //элементы меняются местами по передниму и задниму индексу
	}
	return string(runes) //возваращает строку
}

func main() {
	s := "abcde"            //создаем переменную с типом string
	fmt.Println(Reverse(s)) //запуск функции
}
