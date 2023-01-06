package main

import (
	"fmt"
	"sync"
)

func main() {

	var score = []int{0}

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	wg.Add(1)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {

		fmt.Println("one")
		mu.Lock()
		score = append(score, 1)
		mu.Unlock()
		wg.Done()
	}(wg, mu)
	wg.Add(1)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		fmt.Println("two")
		mu.Lock()
		score = append(score, 2)
		mu.Unlock()
		wg.Done()
	}(wg, mu)
	wg.Add(1)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		fmt.Println("three")
		mu.Lock()
		score = append(score, 3)
		mu.Unlock()
		wg.Done()
	}(wg, mu)

	wg.Wait()
	fmt.Println(score)
}
