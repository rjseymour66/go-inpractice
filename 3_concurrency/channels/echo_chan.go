package main

import (
	"fmt"
	"os"
	"time"
)
/*

ch<-: send to or write to channels. If the arrow
      is on the right, variables send (or write) to the 
      channel.

<-ch: receive from or read from channels. If the 
      arrow is on the left, variables receive (or read) 
      from the channel. 

      You don't have to write to a variable -- you can 
      just write to the program.
*/

func main() {

	done := time.After(30 * time.Second)
	echo := make(chan []byte)
	go readStdin(echo)
	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Timed out")
			os.Exit(0)
		}
	}
}

func readStdin(out chan<- []byte) {
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data
		}
	}
}
