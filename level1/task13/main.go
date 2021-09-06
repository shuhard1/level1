package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) { ////Ошибка в том, что в функцию передается не ссылка на структуру, а копия структуры
			fmt.Println(i)
			wg.Done() //тоесть wg.Done() уменьшает внутренний счетчик активных элементов на единицу в копии
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
