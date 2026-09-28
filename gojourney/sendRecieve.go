package main

import "fmt"

func calculateSquare(num int, ch chan int) {
	result := num * num
	ch <- result // 1. Sending data into the channel
}

func main() {
	ch := make(chan int)

	go calculateSquare(5, ch)

	// 2. Receiving data out of the channel.
	// This automatically blocks 'main' from exiting until the data arrives!
	result := <-ch

	fmt.Println("The result is:", result)
}
