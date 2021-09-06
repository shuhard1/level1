package main

import (
	"fmt"
	"sync"
)

func worker(id int, number <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()       //сигнал, что элемент группы завершил свое выполнение
	fmt.Println(<-number) //чтение <-number и вывод в консоль
}

func main() {
	var wg sync.WaitGroup                                                                    //создаем структуру sync.WaitGroup
	numbers := [20]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19} //создается массив
	number := make(chan int, len(numbers))                                                   //создаются буферизованные каналы
	wg.Add(len(numbers))                                                                     //добавляет N гуротин одну группу

	for w := 1; w <= len(numbers); w++ {
		go worker(w, number, &wg) //запуск горутин
	}

	for j := 0; j <= len(numbers)-1; j++ { //задания отправляются в канал jobs
		number <- numbers[j] //отправка в канал number
	}
	close(number) //закрывает

	wg.Wait() //ожидаем завершения горутин
}
