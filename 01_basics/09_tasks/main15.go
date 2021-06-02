/*


Из натурального числа удалить заданную цифру.

Входные данные

Вводятся натуральное число и цифра, которую нужно удалить.

Выходные данные

Вывести число без заданных цифр.

Sample Input:

38012732
3

Sample Output:

801272

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var N string
	var k int
	fmt.Scan(&N)
	fmt.Scan(&k)
	newN := make([]string, 0)
	for _, i := range N {
		//fmt.Printf("i=%d, k=%d\n", i, k)
		if int(i-'0') != k {
			newN = append(newN, string(i))
		}
	}
	fmt.Print(strings.Join(newN, ""))
}
