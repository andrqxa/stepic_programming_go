/*
Заданы три числа - a, b, c (a < b < c)a,b,c(a<b<c) - длины сторон треугольника. Нужно проверить, является ли треугольник прямоугольным. Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"

Sample Input:

6 8 10
Sample Output:

Прямоугольный
*/
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	a2, b2, c2 := a*a, b*b, c*c
	if a2 == b2+c2 || b2 == a2+c2 || c2 == a2+b2 {
		fmt.Print("Прямоугольный")
	} else {
		fmt.Print("Непрямоугольный")
	}

}
