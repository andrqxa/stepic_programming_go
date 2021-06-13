/*
Необходимо написать функцию func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).

Описание ее работы:

n раз сделать следующее

прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
вычислить f(x1) + f(x2)
записать полученное значение в out
Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.

Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.



Формат ввода:

количество итераций передается через аргумент n.
целые числа подаются через аргументы-каналы in1 и in2.
функция для обработки чисел перед сложением передается через аргумент fn.
Формат вывода:

канал для вывода результатов передается через аргумент out.
*/

package main

import (
	"fmt" // пакет используется для проверки выполнения условия задачи, не удаляйте его
	"math/rand"
	"sync"
	"time" // пакет используется для проверки выполнения условия задачи, не удаляйте его
)

const N = 20

func main() {

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}
	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

//type val map[int]int // order -> value

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	wg1 := &sync.WaitGroup{}
	wg1.Add(n)
	wg2 := &sync.WaitGroup{}
	wg2.Add(n)
	mu1 := &sync.Mutex{}
	mu2 := &sync.Mutex{}
	res1 := make(map[int]int)
	res2 := make(map[int]int)
	idx1 := 0
	idx2 := 0

	go func(src <-chan int) {
		for i1 := range src {
			go func(id int, curr int) {
				fn1 := fn(curr)
				mu1.Lock()
				res1[id] = fn1
				mu1.Unlock()
			}(idx1, i1)
			wg1.Done()
			idx1++
		}
	}(in1)

	go func(src <-chan int) {
		for i2 := range src {
			go func(id int, curr int) {
				fn2 := fn(curr)
				mu2.Lock()
				res2[id] = fn2
				mu2.Unlock()
			}(idx2, i2)
			wg2.Done()
			idx2++
		}
	}(in2)

	go func(out1 chan<- int) {
		wg1.Wait()
		wg2.Wait()
		for i := 0; i < n; i++ {
			out1 <- res1[i] + res2[i]
		}
		// close(out1)
	}(out)
	// fmt.Println(sq)
}
