package main

import (
	"fmt"
	"time"
)

func doble(n int) int { return n * 2 }

func main() {
	dobles := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(500 * time.Millisecond)
			dobles <- doble(i)
		}

		close(dobles)
	}()

	for {
		i, ok := <-dobles
		if !ok {
			break
		}
		fmt.Printf("%d", i)
	}

}
