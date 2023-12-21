package main

import (
	"fmt"
	"time"
)

// 1. 协程工作函数
func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

// 创建的channel只能发数据
func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n",
				id, <-c)
		}
	}()
	return c
}

func chanDemo() {
	// c := make(chan int)
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		// // #1
		// channels[i] = make(chan int)
		// go worker(i, channels[i])
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Microsecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Microsecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()

	fmt.Println("Buffered channel")
	bufferedChannel()

	fmt.Println("Channel close and range")
	channelClose()
}
