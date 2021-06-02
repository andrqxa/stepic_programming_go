/*
Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.

Входные данные

На вход программе подается целое число d (0 < d < 360).

Выходные данные

Выведите на экран фразу:

It is ... hours ... minutes.

Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.

Sample Input:

90
Sample Output:

It is 3 hours 0 minutes.
*/
package main

import (
	"fmt"
)

const (
	template = "It is %d hours %d minutes."
)

func main() {
	var a int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	hr := a / 30
	min := 2 * (a % 30)
	fmt.Printf(template, hr, min)
}
