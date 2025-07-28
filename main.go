package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go func () {
		firstSum := <-c
		fmt.Println(firstSum)
		secondSum := firstSum + 10
		c<-secondSum
	}()

	c <- 100
	fmt.Println(<-c)

}