package main

var justString string

func someFunc() { //Аллокация памяти очень много и очень часто
	//Аллокация -это выделение блока памяти
	//Сборщик мусора будет тормозить пытаясь собрать неиспользуюемую память
	v := createHugeString(1 << 10) //1 умножается на 2, 10 раз
	justString = v[:100]
}
func main() {
	someFunc()
}

//Решение: Не объявлять в функции большую строку, объявить ее снаружи и подавать в функцию по ссылке
