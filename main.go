package main

import "sync"

func expensiveOp() {
	sum := 0
	for i := range 1000000 {
		sum += i
	}
	println(sum)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			expensiveOp()
		}()
	}
	wg.Wait()
	println("main Function ended")
}
