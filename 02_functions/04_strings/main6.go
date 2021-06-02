/*


Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input:

fdsghdfgjsdDD1

Sample Output:

Ok

*/

package main

import (
	"fmt"
	"regexp"
)

const templ = "^[a-zA-Z0-9]{5,}$"

func main() {
	var srcText string
	fmt.Scan(&srcText)
	match, _ := regexp.MatchString(templ, srcText)
	if match {
		fmt.Print("Ok")
		return
	}
	fmt.Print("Wrong password")
}
