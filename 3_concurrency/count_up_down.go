package main

import (
	"fmt"
	"time"
)

func main() {
	go counter()
	time.Sleep(15 * time.Second)
	fmt.Println("Done")
}

func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}

	for i := 10; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}
