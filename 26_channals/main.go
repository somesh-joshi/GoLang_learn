package main

import (
	"fmt"
	"sync"
)

func main() {

	myCha := make(chan int, 10)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		val, ok := <-ch
		fmt.Println(val, ok)

		wg.Done()
	}(myCha, wg)

	go func(ch chan int, wg *sync.WaitGroup) {

		close(myCha)
		wg.Done()
	}(myCha, wg)

	wg.Wait()
}
