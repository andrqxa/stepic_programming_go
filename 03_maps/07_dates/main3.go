/*
На стандартный ввод подается строковое представление двух дат, разделенных
запятой (формат данных смотрите в примере).

Необходимо преобразовать полученные данные в тип Time, а затем вывести
продолжительность периода между меньшей и большей датами.

Sample Input:

13.03.2018 14:00:15,12.03.2018 14:00:15
Sample Output:

24h0m0s
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	s, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	x := strings.Split(s, ",")
	first, second := x[0], x[1]
	//second = second[:len(second)-1]

	firstTime, err := time.Parse("02.01.2006 15:04:05", first)
	if err != nil {
		panic(err)
	}
	secondTime, err := time.Parse("02.01.2006 15:04:05", second)
	if err != nil {
		panic(err)
	}
	if secondTime.Before(firstTime) {
		firstTime, secondTime = secondTime, firstTime
	}

	duration := secondTime.Sub(firstTime).String()
	fmt.Println(duration)
}
