package main

import (
	"fmt"
)

type void struct{} //создается пустая структура

var member void

//cat, cat, dog, cat, tree

//этот способ сохраняет байты
func main() {
	setLive := make(map[string]void) //создается пустая карта
	setLive["cat"] = member          //добавления во множество элементы
	setLive["dog"] = member
	setLive["tree"] = member

	for k := range setLive { //вывод в консоль
		fmt.Println(k)
	}
}
