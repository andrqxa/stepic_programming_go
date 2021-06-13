package main

import (
	"fmt"
	"sync"
)

type (
	val map[int]int
	seq map[int]val
)

const N = 3

func fn(i int) int {
	return 1 << i
}

func main() {
	sq := make(seq)
	sq[1] = val{0: 0}
	sq[2] = val{0: 0}
	in1 := make(chan int, N)
	wg1 := &sync.WaitGroup{}
	wg1.Add(N)
	mu := &sync.Mutex{}
	go func(src <-chan int) {
		idx := 0
		for i := range src {
			fmt.Printf("From channel 1 <- %d\n", i)
			k := i
			go func(idx1 int) {
				res1 := fn(k)
				mu.Lock()
				sq[1] = val{idx1: res1}
				fmt.Printf("i=%d, f(%d)=%d\n", k, k, res1)
				fmt.Println(sq)
				mu.Unlock()
				wg1.Done()
			}(idx)
			idx++
		}
	}(in1)
	for i := 0; i < N; i++ {
		in1 <- i
	}
	wg1.Wait()
	fmt.Println(sq)
}
