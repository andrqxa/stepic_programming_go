/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные

Вводится четыре числа.

Выходные данные

Необходимо вернуть из функции наименьшее из 4-х данных чисел

Sample Input:

4 5 6 7

Sample Output:

4
*/
func minimumFromFour() int {
	var curr int
	fmt.Scan(&curr)
	min := curr
	for i := 0; i < 3; i++ {
		fmt.Scan(&curr)
		if curr < min {
			min = curr
		}
	}
	return min
}