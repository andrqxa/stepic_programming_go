/*
Количество нулей
По данным числам, определите количество чисел, которые равны нулю.

Входные данные
Вводится натуральное число N, а затем N чисел.

Выходные данные
Выведите количество чисел, которые равны нулю.

Sample Input:

5
1
8
100
0
12
Sample Output:

1
*/
package main

import "fmt"

func main() {
	var N, curr int
	count := 0
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&curr)
		if curr == 0 {
			count++
		}
	}
	fmt.Print(count)
}
