package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 25; i++ {
		time.Sleep(time.Minute)
		fmt.Printf("WORKING %d/25\n", i)
	}

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Minute)
		fmt.Printf("RESTING %d/5\n", i)
	}

}
