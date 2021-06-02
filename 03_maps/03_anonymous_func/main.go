/*


Используем анонимные функции на практике.

Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint, которую внутри функции main в дальнейшем можно будет вызвать по имени fn.

Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные цифры и цифра 0. Если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

Пакет main объявлять не нужно!
Вводить и выводить что-либо не нужно!
Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.

Sample Input:

727178

Sample Output:

28
*/

// fn
fn := func(src uint) uint {
	if src == 0 {
		return uint(100)
	}
	str := strconv.Itoa(int(src))
	var res string
	for _, r := range str {
		intR, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		if intR == 0 || intR%2 == 1 {
			continue
		}
		res += string(r)
	}
	if len(res) == 0 {
		return uint(100)
	}
	intRes, err := strconv.Atoi(string(res))
	if err != nil {
		panic(err)
	}
	return uint(intRes)
}