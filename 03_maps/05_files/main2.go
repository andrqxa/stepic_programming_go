/*
Данная задача поможет вам разобраться в пакете encoding/csv и path/filepath,
хотя для решения может быть использован также пакет archive/zip (поскольку файл
с заданием предоставляется именно в этом формате).

В тестовом архиве, который вы можете скачать из нашего репозитория на github.com,
содержится набор папок и файлов. Один из этих файлов является файлом с данными в
формате CSV, прочие же файлы структурированных данных не содержат.

Требуется найти и прочитать этот единственный файл со структурированными данными
(это таблица 10х10, разделителем является запятая), а в качестве ответа
необходимо указать число, находящееся на 5 строке и 3 позиции (индексы 4 и 2
соответственно).
*/

package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	const root = "./task"
	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}

	if info.IsDir() {
		return nil // Проигнорируем директории
	}
	bb, errr := ioutil.ReadFile(path)
	if errr != nil {
		return fmt.Errorf("open file  error [%v]\n", errr)
	}
	if strings.Contains(string(bb), ",") {
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("open file  error [%v]\n", err)
		}
		defer file.Close()

		r := csv.NewReader(file)
		data, err := r.ReadAll()
		if err != nil {
			//
		}
		fmt.Println(data[4][2])
	}
	return nil
}
