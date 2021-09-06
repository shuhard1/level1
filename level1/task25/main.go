package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *Counter) Inc(key string) { // увеличивает счетчик для данного ключа
	defer c.mu.Unlock() //разблокировка.defer - выполняется самым последним в функции
	c.mu.Lock()         //заблокрировать, чтобы только одна горутина могла получить доступ
	c.v[key]++          //инкрементирование
	w := c.v[key]       //переменной w присваивается значение c.v[key]
	fmt.Println(w)      //вывод w в консоль
}

func main() {
	c := Counter{v: make(map[string]int)} //создаем карту
	for i := 0; i < 1000; i++ {           //запуск 1000 гуротин
		go c.Inc("somekey")
	}

	time.Sleep(time.Second) //даем время, чтобы main не завершил свою работу и горутины успели
}
