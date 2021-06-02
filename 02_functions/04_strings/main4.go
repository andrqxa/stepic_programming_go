/*
На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

Sample Input:

ihgewlqlkot

Sample Output:

hello
*/

package main

import "fmt"

func main() {
	var srcText string
	fmt.Scan(&srcText)
	src := []rune(srcText)
	for i, r := range src {
		if i%2 != 0 {
			fmt.Print(string(r))
		}
	}
}
