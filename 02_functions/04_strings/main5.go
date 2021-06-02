/*
Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку

Sample Input:

zaabcbd

Sample Output:

zcd
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var srcText string
	fmt.Scan(&srcText)
	src := []rune(srcText)
	for _, v := range src {
		if strings.Count(srcText, string(v)) > 1 {
			srcText = strings.ReplaceAll(srcText, string(v), "")
		}
	}
	fmt.Print(srcText)
}
