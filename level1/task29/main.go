package main

import (
	"fmt"
	"math/big"
)

func main() {
	var res big.Int
	a := big.NewInt(6689348814741910300) //создаем переменную a с типом *big.Int
	b := big.NewInt(368934881474191030)  ////создаем переменную a с типом *big.Int

	fmt.Print("Умножение: ")
	fmt.Println(res.Mul(a, b)) //Mul - умножает a на b и вывод в консоль

	fmt.Print("Деление: ")
	fmt.Println(res.Div(a, b)) //Div - делит a на b и вывод в консоль

	fmt.Print("Сложение: ")
	fmt.Println(res.Add(a, b)) //Add - складывает a и b и вывод в консоль

	fmt.Print("Вычетание: ")
	fmt.Println(res.Sub(a, b)) ////Sub - вычетает b из a и вывод в консоль

}
