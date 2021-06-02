/*
Дано трехзначное число. Переверните его, а затем выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных
Выведите перевернутое число.

Sample Input:

843
Sample Output:

348
*/
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a, b, c := n/100, n/10%10, n%10
	fmt.Printf("%d%d%d", c, b, a)
}
