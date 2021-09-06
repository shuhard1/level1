package main

import (
	"fmt"
	"math"
)

// Точка представляет собой точку в пространстве.
type Point struct {
	X int
	Y int
}

// New возвращает точку на основе позиций X и Y на графике.
func New(x int, y int) Point {
	return Point{x, y}
}

// Distance находит длину гипотенузы между двумя точками.
func (p Point) Distance(p2 Point) float64 {
	first := math.Pow(float64(p2.X-p.X), 2)  //math.Pow возводит float64(p2.X-p.X) в степень 2
	second := math.Pow(float64(p2.Y-p.Y), 2) ///math.Pow возводит float64(p2.X-p.Y) в степень 2
	return math.Sqrt(first + second)         //math.Sqrt возвращает кубический корень из результата сложения first + second
}

func main() {
	p1 := New(37, -76) //создаем переменную с типом Point
	p2 := New(26, -66) //создаем переменную с типом Point

	dist := p1.Distance(p2)       //запускаем функцию
	fmt.Println("Distance", dist) //вывод в консоль dist
}
