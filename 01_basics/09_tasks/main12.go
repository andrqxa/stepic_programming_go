/*


По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

Входные данные

Вводится натуральное число.

Выходные данные

Выведите ответ на задачу.

Sample Input:

50

Sample Output:

1 2 4 8 16 32


*/
package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	cur := 1
	for cur < N+1 {
		fmt.Printf("%d ", cur)
		cur *= 2
	}
}
