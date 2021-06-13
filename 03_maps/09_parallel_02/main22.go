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
	ch := make(chan int, N)
	wg1 := &sync.WaitGroup{}
	wg1.Add(N)
	mu := &sync.Mutex{}
	// sq[2] = val{0: 0}
	idx := 0
	go func(c chan int) {
		// var curr int
		for i := range c {
			go func(k int) {
				curr := fn(k)
				mu.Lock()
				rs := sq[1]
				rs[idx] = curr
				sq[1] = rs
				idx++
				mu.Unlock()
				wg1.Done()
			}(i)

		}

	}(ch)
	for i := 0; i < N; i++ {
		ch <- i
	}
	wg1.Wait()
	close(ch)
	fmt.Println(sq)
}
