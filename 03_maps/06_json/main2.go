/*
Данная задача ориентирована на реальную работу с данными в формате JSON. Для
решения мы будем использовать справочник ОКВЭД (Общероссийский классификатор
видов экономической деятельности), который был размещен на web-портале
data.gov.ru.

Необходимая вам информация о структуре данных содержится в файле
structure-20190514T0000.json, а сами данные, которые вам потребуется
декодировать, содержатся в файле data-20190514T0100.json. Файлы размещены в
нашем репозитории на github.com.

Для того, чтобы показать, что вы действительно смогли декодировать документ вам
необходимо в качестве ответа записать сумму полей global_id всех элементов,
закодированных в файле.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	dataFile = "data-20190514T0100.json"
	// jsonFile = "structure-20190514T0000.json"
)

type Property struct {
	System_object_id string `json:"system_object_id"`
	Kod              string
	Signature_date   string `json:"signature_date"`
	Global_id        int    `json:"global_id"`
	Idx              string
	Razdel           string
	Name             string
}

type Data []Property

func main() {
	jsonFile, err := ioutil.ReadFile(dataFile)
	if err != nil {
		panic(err)
	}
	var data Data
	if err := json.Unmarshal(jsonFile, &data); err != nil {
		fmt.Println(err)
		return
	}
	res := 0
	for _, v := range data {
		res += v.Global_id
	}
	fmt.Printf("%d", res)
}
