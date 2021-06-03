package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(os.Stdout)
	sum := 0

	for scanner.Scan() {
		curr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("invalid arguments")
		}
		sum += curr
	}

	res := strconv.Itoa(sum)

	// fmt.Printf("sum = %d\n", sum)
	// fmt.Printf("sum1 = %s %T", res, res)

	_, err = writer.WriteString(res)
	if err != nil {
		panic("can't write to console")
	}
	// fmt.Printf("Записано %d байт\n", n) // Записано 27 байт
	writer.Flush()
}
