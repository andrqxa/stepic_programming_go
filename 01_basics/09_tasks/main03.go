/*
Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток. Например, если

k=13257=3*3600+40*60+57,

то h=3 и m=40.

Входные данные

На вход программе подается целое число k (0 < k < 86399).

Выходные данные

Выведите на экран фразу:

It is ... hours ... minutes.

Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.

Sample Input:

13257
Sample Output:

It is 3 hours 40 minutes.
*/
package main

import "fmt"

const templ = "It is %d hours %d minutes."

func main() {
	var n int
	fmt.Scan(&n)
	hr, min := n/3600, n%3600/60
	fmt.Printf(templ, hr, min)
}
