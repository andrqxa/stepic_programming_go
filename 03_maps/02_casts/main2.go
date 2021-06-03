/*


Представьте что вы работаете в большой компании где используется модульная архитектура. Ваш коллега написал модуль с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные. Вы же пишете функцию которая считывает две переменных типа string ,  а возвращает число типа int64 которое нужно получить сложением этих строк.


Но не было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того что гоферам платят больше. Поэтому он решил подшутить над вами и подсунул вам подвох. Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.


Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа. Под мусором имеются ввиду лишние символы и спец знаки. Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes. Они уже импортированы, вам ничего импортировать не нужно!

Считывать и выводить ничего не нужно!

Ваша функция должна называться adding() !

Считайте что функция и пакет main уже объявлены!

Sample Input:

%^80 hhhhh20&&&&nd

Sample Output:

100

*/
func adding(a, b string) int64 {
	srcA := []rune(a)
	srcB := []rune(b)
	clearA := make([]int, 0)
	clearB := make([]int, 0)
	for _, r := range srcA {
		dig := r - '0'
		if dig >= 0 && dig <= 9 {
			clearA = append(clearA, int(dig))
		}
	}
	for _, r := range srcB {
		dig := r - '0'
		if dig >= 0 && dig <= 9 {
			clearB = append(clearB, int(dig))
		}
	}
	var strA, strB string
	for _, i := range clearA {
		strA += strconv.Itoa(i)
	}
	for _, i := range clearB {
		strB += strconv.Itoa(i)
	}
	resA, _ := strconv.Atoi(strA)
	resB, _ := strconv.Atoi(strB)
	return int64(resA + resB)
}