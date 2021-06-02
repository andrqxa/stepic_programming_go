/*


По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".

Входные данные

Дано число n (0<n<100).

Выходные данные

Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.

Sample Input:

10

Sample Output:

10 korov


*/
package main

import (
	"fmt"
)

const templ = "%d korov%s"

func main() {
	var kor int
	var odd string
	fmt.Scan(&kor)
	if kor > 10 && kor < 20 {
		fmt.Printf(templ, kor, "")
		return
	}
	oddKor := kor % 10
	switch oddKor {
	case 1:
		odd = "a"
	case 2, 3, 4:
		odd = "y"
	default:
		odd = ""
	}
	fmt.Printf(templ, kor, odd)
}
