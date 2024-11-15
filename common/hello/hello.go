package main

import (
	"fmt"
	"sync"
)

func main() {
	hello()
}

func hello() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			do()
		}()
	}
	wg.Wait()
	fmt.Println("finished")
}

func do() {
	fmt.Println("do panic")
	panic("panic")
}
