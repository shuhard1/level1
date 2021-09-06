package main

import (
	"fmt"
	"sync"
)

func main() {
	var numbers = [5]int{2, 4, 6, 8, 10} //создаем массив
	var wg sync.WaitGroup                //определяем группу в виде переменной wg sync.WaitGroup

	wg.Add(len(numbers)) //число, которое передается в метод Add определяет значение внутреннего счетчика активных элементов.

	for _, number := range numbers { //проходимся по массиву
		go CalcSquare(&wg, number) //запускаем горутины
	}

	wg.Wait() // ожидаем завершения горутин
	fmt.Println("Горутины завершили выполнение")
}

func CalcSquare(wg *sync.WaitGroup, number int) {
	defer wg.Done() //сигнал, что элемент группы завершил свое выполнение. defer - выполянется самым последним в функции
	//Горутина начала выполнение
	number *= number    //возведение number в квадрат
	fmt.Println(number) //вывод в консоль
	//Горутина завершила выполнение
}
