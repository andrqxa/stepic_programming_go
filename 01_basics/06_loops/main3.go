/*
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8. Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.

Sample Input:

5
38 24 800 8 16
Sample Output:

40
*/
package main

import "fmt"

func main() {
	var leng, next int
	fmt.Scan(&leng)
	sum := 0
	for i := 0; i < leng; i++ {
		fmt.Scan(&next)
		if next > 9 && next < 100 && next%8 == 0 {
			sum += next
		}
	}
	fmt.Println(sum)
}
