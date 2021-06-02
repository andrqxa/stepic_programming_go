/*
Количество минимумов
Найдите количество минимальных элементов в последовательности.



Входные данные

Вводится натуральное число N, а затем N чисел.

Выходные данные

Выведите количество минимальных элементов.

Sample Input:

3
21
11
4
Sample Output:

1
*/
package main

import "fmt"

func main() {
	var N, curr int
	arr := make([]int, 0)
	fmt.Scan(&N)
	fmt.Scan(&curr)
	arr = append(arr, curr)
	currMin := curr
	for i := 1; i < N; i++ {
		fmt.Scan(&curr)
		arr = append(arr, curr)
		if curr < currMin {
			currMin = curr
		}
	}
	count := 0
	for _, i := range arr {
		if i == currMin {
			count++
		}
	}
	//fmt.Println(arr)
	//fmt.Println(currMin)
	fmt.Println(count)
}
