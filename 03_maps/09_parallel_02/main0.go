package main

import (
	"sync"
)

const N = 3

// import "fmt"

// func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
func merge2Channels(in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	mapRes := make(map[int][2]int)
	wg1 := &sync.WaitGroup{}
	wg1.Add(N)
	wg2 := &sync.WaitGroup{}
	wg2.Add(N)
	mu := &sync.Mutex{}
	go func(src <-chan int) {
		idx := 0
		for i := range src {
			mu.Lock()
			_, ok := mapRes[idx]
			if !ok {
				mapRes[idx] = [2]int{0, 0}
			}
			arr := mapRes[idx]
			arr[0] = i
			mapRes[idx] = arr
			idx++
			mu.Unlock()
			wg1.Done()
		}
	}(in1)
	go func(src <-chan int) {
		idx := 0
		for i := range src {
			mu.Lock()
			_, ok := mapRes[idx]
			if !ok {
				mapRes[idx] = [2]int{0, 0}
			}
			arr := mapRes[idx]
			arr[1] = i
			mapRes[idx] = arr
			idx++
			mu.Unlock()
			wg2.Done()
		}
	}(in2)
	go func(rs chan<- int) {
		wg1.Wait()
		wg2.Wait()
		for _, v := range mapRes {
			out <- v[0] + v[1]
		}
		close(out)
	}(out)

}

func main() {

	// in1 := make(chan int, N)
	// in2 := make(chan int, N)
	// out := make(chan int, N)

	// for i := 0; i < N; i++ {
	// 	in1 <- 3 + i
	// 	in2 <- 7 * i
	// }
	// merge2Channels(in1, in2, out, N)
	// close(in1)
	// close(in2)

	// for v := range out {
	// 	fmt.Println(v)
	// }

}

// func merge2Channels(in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
// }
