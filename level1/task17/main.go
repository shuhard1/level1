package main

import (
	"fmt"
)

var t interface { //создается интерфейс t
	talk() string
}

type martian struct{} // создается пустая структура martian

func main() {
	t = martian{}
	fmt.Println(typeof(t)) //запуск функции и вывод в консоль

}

func typeof(Printer interface{}) string {
	switch t := Printer.(type) {
	case interface{ talk() int }: //если t имеет тип interface{ talk() int },
		return "int" //то вернуть строку "int"
	case interface{ talk() string }: //и т.д.
		return "string"
	case interface{ talk() bool }:
		return "bool"
	case interface{ talk() chan int }:
		return "chan int"
	case interface{ talk() chan bool }:
		return "chan bool"
	case interface{ talk() chan string }:
		return "chan bool"
	default: //если ничего не подошло,
		_ = t
		return "unknown" //то вернуть строку "unknown"
	}
}

func (m martian) talk() string {
	return "nack nack"
}
