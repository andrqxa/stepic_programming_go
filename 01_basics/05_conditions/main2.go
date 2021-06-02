/*
По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".

Sample Input 1:

237
Sample Output 1:

YES
Sample Input 2:

117
Sample Output 2:

NO
*/
package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	hun := a / 100
	ten := a / 10 % 10
	one := a % 10
	// fmt.Printf("%d %d %d\n", hun, ten, one)
	if hun != ten && ten != one && one != hun {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
