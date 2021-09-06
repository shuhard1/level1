package main

import (
	"fmt"
)

//cat, cat, dog, cat, tree
func main() {
	setLive := make(map[string]bool) //создается пустая карта
	setLive["cat"] = true            //добавления во множество элементы
	setLive["dog"] = true
	setLive["tree"] = true

	for k := range setLive { //вывод в консоль
		fmt.Println(k)
	}
}
