/*


На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы

Sample Input:

6 8

Sample Output:

10
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		panic("Invalid arguments")
	}
	if a <= 0 || b <= 0 {
		panic("Arguments must be positive")
	}
	fmt.Print(math.Sqrt(a*a + b*b))
}
