/*

Самое большое число, кратное 7

Найдите самое большее число на отрезке от a до b, кратное 7 .

Входные данные
Вводится два целых числа a и b (a≤b).

Выходные данные
Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b), кратное 7 , или выведите "NO" - если таковых нет.

Sample Input:

100
500

Sample Output:

497


*/
package main

import (
	"fmt"
)

func main() {
	var a, b, max int
	fmt.Scan(&a, &b)
	max = a
	oddCount := 0
	for i := a; i < b+1; i++ {
		odd := i % 7
		if odd == 0 && i > max {
			max = i
			oddCount++
		}
	}
	if oddCount == 0 {
		fmt.Print("NO")
		return
	}
	fmt.Print(max)
}
