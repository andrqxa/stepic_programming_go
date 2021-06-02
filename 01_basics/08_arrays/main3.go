/*
На ввод подаются пять чисел, которые записываются в массив. Однако эта часть программы уже написана.

Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.

Sample Input:

2
3
56
45
21
Sample Output:

56
*/
package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	//fmt.Printf("%v\n", array)
	if len(array) == 0 {
		panic("array is empty")
	}
	if len(array) == 1 {
		fmt.Print(array[0])
		return
	}
	max := array[0]
	for _, v := range array[1:] {
		if v >= max {
			max = v
		}
	}
	fmt.Print(max)
}
