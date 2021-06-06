/*
На стандартный ввод подаются данные о студентах университетской группы в формате JSON:

{
    "ID":134,
    "Number":"ИЛМ-1274",
    "Year":2,
    "Students":[
        {
            "LastName":"Вещий",
            "FirstName":"Лифон",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[1,2,3]
        },
        {
            // ...
        }
    ]
}
В сведениях о каждом студенте содержится информация о полученных им оценках
(Rating). Требуется прочитать данные, и рассчитать среднее количество оценок,
полученное студентами группы. Ответ на задачу требуется записать на стандартный
вывод в формате JSON в следующей форме:

{
    "Average": 14.1
}
Как вы понимаете, для декодирования используется функция Unmarshal, а для
кодирования MarshalIndent (префикс - пустая строка, отступ - 4 пробела).

Если у вас возникли проблемы с чтением / записью данных, то этот комментарий для
вас: в уроках об интерфейсах и работе с файлами мы рассказывали, что стандартный
ввод / вывод - это файлы, к которым можно обратиться через os.Stdin и os.Stdout
соответственно, они удовлетворяют интерфейсам io.Reader и io.Writer, из них
можно читать, в них можно писать.

Один из способов чтения, использовать ioutil.ReadAll:

data, err := ioutil.ReadAll(os.Stdin)
if err != nil {
    // ...
}

// data - тип []byte
Задачу можно выполнить и другими способами, в частности использовать bufio.
Буквально в следующем шаге (через один, на самом деле) будет рассказано о еще
одном способе чтения / записи, можете забежать немного вперед, а затем вернуться
к задаче.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

func (s Student) StudentCountRating() int {
	return len(s.Rating)
}

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

func (g *Group) GroupAverageCount() (res float64) {
	sum := 0
	for _, r := range g.Students {
		sum += r.StudentCountRating()
	}
	res = float64(sum) / float64(len(g.Students))
	return
}

type ResAverage struct {
	Average float64
}

func PrintResult(res interface{}) {
	switch res.(type) {
	case ResAverage:
		{
			data, err := json.MarshalIndent(res, "", "    ")
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", data)
		}
	default:
		panic("Invalid json")
	}

}

func main() {
	jsonFile, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	var g Group
	if err := json.Unmarshal(jsonFile, &g); err != nil {
		fmt.Println(err)
		return
	}
	avg := ResAverage{g.GroupAverageCount()}

	PrintResult(avg)

}
