package main

import (
	"fmt"
	//"runtime"
)

func fibonacci(c, quit chan int) {
	fmt.Println("t2")
	x, y := 1, 1
	for {
		fmt.Println("bb")
		select {
		case c <- x:
			fmt.Println("aa")
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	//runtime.GOMAXPROCS(1)
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("t1")
		for i := 0; i < 3; i++ {
			fmt.Println("t3")
			n := <-c
			fmt.Println("t4")
			fmt.Println(n)
		}
		quit <- 0
	}()

	go fibonacci(c, quit)

	for {
	}
}
