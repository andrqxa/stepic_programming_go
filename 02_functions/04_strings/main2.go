/*


На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является паллиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").

Sample Input:

топот

Sample Output:

Палиндром

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func trimAll(sourse []rune) []rune {
	res := make([]rune, 0)
	for _, r := range sourse {
		if !unicode.IsSpace(r) {
			res = append(res, r)
		}
	}
	return res
}

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// text := "ааиа" //"А роза упала на лапу Азора"
	text = strings.ToLower(text)
	runes := trimAll([]rune(text))
	for i, r := range runes[:len(runes)/2] {
		if r != runes[len(runes)-1-i] {
			fmt.Println("Нет")
			return
		}
	}
	fmt.Println("Палиндром")
}
