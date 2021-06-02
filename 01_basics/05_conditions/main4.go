/*
Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.



Формат входных данных

На вход подается номер билета - одно шестизначное  число.

Формат выходных данных
Выведите "YES", если билет счастливый, в противном случае - "NO".

Sample Input:

613244
Sample Output:

YES
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	fmt.Scan(&a)
	first := a[:3]
	second := a[3:]
	if isEqual(first, second) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func isEqual(a, b string) bool {
	return sum(a) == sum(b)
}

func sum(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	one := i / 100
	two := i / 10 % 10
	three := i % 10
	return one + two + three
}
