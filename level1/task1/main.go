package main

import (
	"fmt"
)

func main() {
	Man := Human{"Mark"}             //создается переменная с типом Human
	action := Action{Man, "nothing"} //создается переменная с типом Action
	fmt.Println(action)              //вывод в консоль
}

type Human struct { //родительской структура
	Name string
}

type Action struct { //дочерняя струткура
	Human
	Do string
}
