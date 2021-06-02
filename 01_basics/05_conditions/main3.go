/*
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

Формат входных данных
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.

Sample Input:

1234
Sample Output:

1
*/
package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a >= 0 && a < 10:
		fmt.Print(a)
	case a >= 10 && a < 100:
		fmt.Print(a / 10)
	case a >= 100 && a < 1000:
		fmt.Print(a / 100)
	case a >= 1000 && a < 10000:
		fmt.Print(a / 1000)
	case a >= 10000:
		fmt.Print(a / 10000)
	}
}
