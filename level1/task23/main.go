package main

import "fmt"

func main() {
	searchField := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91} //создаем уже отсортированный массив срез
	searchNumber := 91                                        //сохдаем искомое число

	fmt.Println("Running Program")                          //вывод в консоль
	fmt.Println("Searching list of numbers: ", searchField) //вывод в консоль срез
	fmt.Println("Searching for number: ", searchNumber)     //вывод в консоль число, которое надо найти

	result := binarySearch(searchField, searchNumber)                //бинарный поиск
	fmt.Println("Found! Your number is found in position: ", result) //вывод результата функции в консоль
}

func binarySearch(a []int, search int) (result int) {
	mid := len(a) / 2 //делим длину среза пополам, чтобы начать проверку с середины
	switch {
	case len(a) == 0: //если длина среза равна 0
		result = -1 //то функция вернет result = -1
	case a[mid] > search: //если число в центре среза больше искомого,
		result = binarySearch(a[:mid], search) //то ищем число влево этим же методом отбрасывая половину среза
		fmt.Println(mid)                       //вывод в консоль для наглядности
	case a[mid] < search: //если число в центре среза меньше искомого,
		result = binarySearch(a[mid+1:], search) //то ищем число справо этим же методом оотбрасывая половину среза
		fmt.Println(mid)
		if result >= 0 { // если что-либо, кроме результата -1 не найдено
			result += mid + 1
		}
	default: // a[mid] равно искомому числу, то
		result = mid //позиция числа найдена
		fmt.Println(mid)
	}
	return
}
