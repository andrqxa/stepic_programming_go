/*

Двоичная запись

Дано натуральное число N. Выведите его представление в двоичном виде.

Входные данные

Задано единственное число N

Выходные данные

Необходимо вывести требуемое представление числа N.

Sample Input:

12

Sample Output:

1100
*/
package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	res := make([]int, 0)
	for N/2 > 0 {
		res = append(res, N%2)
		N = N / 2
	}
	res = append(res, N)
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i])
	}
}
