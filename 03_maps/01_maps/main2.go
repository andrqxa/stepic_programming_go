/*


В ходе анализа результатов переписи населения информация была сохранена в объекте типа map:

groupCity := map[int][]string{
	10:   []string{...}, // города с населением 10-99 тыс. человек
	100:  []string{...}, // города с населением 100-999 тыс. человек
	1000: []string{...}, // города с населением 1000 тыс. человек и более
}

При подготовке важного отчета о городах с населением 100-999 тыс. человек был подготовлен другой объект типа map:

cityPopulation := map[string]int{
	город: население города в тысячах человек,
}

Однако база данных с информацией о точной численности населения содержала ошибки, поэтому в cityPopulation в т.ч. была сохранена информация о городах, которые входят в другие группы из groupCity.

Ваша программа имеет доступ к обоим указанным отображениям, требуется исправить cityPopulation, чтобы в ней была сохранена информация только о городах из группы groupCity[100].

Функция main() уже объявлена, доступ к отображениям осуществляется по указанным именам. По результатам выполнения ничего больше делать не требуется, проверка будет осуществлена автоматически.

*/

/*
 * Группировка городов по численности населения в тысячах человек
 * от 10 до 100, от 100 до 1000 и более 1000:
 * groupCity map[int][]string{
 *	 10: []string{...},
 *	 100: []string{...},
 *	 1000: []string{...},
 * }
 *
 * Население городов в тысячах человек:
 * cityPopulation map[string]int{...}
 */
city100 := make(map[string]int, 0)
for _, city := range groupCity[100] {
	city100[city] = 1
}
for city, _ := range cityPopulation {
	if _, ok := city100[city]; !ok {
		delete(cityPopulation, city)
	}
}