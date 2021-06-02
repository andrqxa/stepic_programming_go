/*


Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.

Входные данные

Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только десятичные цифры.

Выходные данные

Выведите максимальную цифру, которая встречается во введенной строке.

Sample Input:

1112221112

Sample Output:

2

*/
package main

import (
	"fmt"
)

func main() {
	var a string
	_, err := fmt.Scan(&a)
	if err != nil {
		panic("Invalid arguments")
	}
	src := []rune(a)
	max := src[0]
	for _, s := range src[1:] {
		if s == '9' {
			max = s
			break
		}
		if s > max {
			max = s
		}
	}
	fmt.Print(string(max))
}
