/*
Поэтапный поиск данных
Данная задача в основном ориентирована на изучение типа bufio.Reader, поскольку
этот тип позволяет считывать данные постепенно.

В тестовом файле, который вы можете скачать из нашего репозитория на github.com,
содержится длинный ряд чисел, разделенных символом ";". Требуется найти, на
какой позиции находится число 0 и указать её в качестве ответа. Требуется
вывести именно позицию числа, а не индекс (то-есть порядковый номер, нумерация
с 1).

Например:  12;234;6;0;78 :
Правильный ответ будет порядковый номер числа: 4
*/
package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

const (
	fileName = "task.data"
	// templ    = "([\\d;]+);0;([\\d;]+)"
)

func main() {
	// reg := regexp.MustCompile(";0;")
	bb, _ := ioutil.ReadFile(fileName)
	// if err != nil {
	// 	panic("open file  error")
	// }
	// fistHalf := reg.Split(string(bb), 2)[0]
	// res := len(strings.Split(fistHalf, ";")) + 1
	// fmt.Println(res)

	fmt.Println(len(strings.Split(strings.Split(string(bb), ";0;")[0], ";")) + 1)
	fmt.Println(len(strings.Split(regexp.MustCompile(";0;").Split(string(bb), 2)[0], ";")) + 1)

	// fmt.Println(len(strings.Split(regexp.MustCompile(";0;").Split(string(bb), 2)[0], ";")) + 1)
}
