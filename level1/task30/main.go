package main

import (
	"fmt"
)

func main() {
	a := []string{"A", "B", "C", "D", "E"} //Создать срез
	i := 1                                 // Индекс элемента, которого надо удалить

	// Удалить элемент по индексу i из a.

	// 1. Выполнить сдвиг a[i+1:] влево на один индекс.
	copy(a[i:], a[i+1:])

	// 2. Удалить последний элемент (записать нулевое значение).
	a[len(a)-1] = ""

	// 3. Усечь срез.
	a = a[:len(a)-1]

	fmt.Println(a) // [A B D E]
}
