/*

Номер числа Фибоначчи

Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.

Входные данные

Вводится натуральное число.

Выходные данные

Выведите ответ на задачу.

Sample Input:

8

Sample Output:

6


*/
package main

import (
	"fmt"
)

func fib(a int) int {
	first, second := 1, 1
	cnt := 2
	for second <= a {
		first, second = second, first+second
		//fmt.Printf("first=%d, second=%d\n", first, second)
		cnt++
	}

	// fmt.Printf("a=%d, first=%d, second=%d, count=%d\n", a, first, second, cnt)
	if first != a {
		return -1
	}
	return cnt - 1
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Print(fib(N))

}
