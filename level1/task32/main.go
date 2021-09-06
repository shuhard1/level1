package main

import (
	"fmt"
	"time"
)

func main() {
	Sleep(5)
	fmt.Println("programme finished") //вывод в консоль
}

func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x)) //возвращает вам канал. И значение будет отправлено по каналу после указанной продолжительности
}
