/*
Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0. Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).

Входные данные

Сначала задано число NN — количество элементов в массиве (1 \leq N \leq 1001≤N≤100). Далее через пробел записаны NN чисел — элементы массива. Массив состоит из целых чисел.

Выходные данные

Необходимо вывести все элементы массива с чётными номерами.

Sample Input:

6
4 5 3 4 2 3
Sample Output:

4 3 2
*/
package main

import "fmt"

func main() {
	odd := func(a int) bool {
		return a%2 == 0
	}
	array := make([]int, 0)
	var ln int
	var curr int
	fmt.Scan(&ln)
	for i := 0; i < ln; i++ {
		fmt.Scan(&curr)
		if odd(i) {
			array = append(array, curr)
		}
	}
	for _, v := range array {
		fmt.Printf("%d ", v)
	}
}
