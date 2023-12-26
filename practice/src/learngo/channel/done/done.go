package main

import (
	"fmt"
	"sync"
)

// 1. 协程工作函数
func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		// go func() {
		// 	done <- true
		// }()
		w.done()
	}
}

type worker struct {
	in chan int
	// done chan bool
	done func()
}

// 创建的channel只能发数据
func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		// done: make(chan bool),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	// c := make(chan int)
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		// // #1
		// channels[i] = make(chan int)
		// go worker(i, channels[i])
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
	// // 接收done chan的完成信号
	// for _, worker := range workers {
	// 	<-worker.done
	// 	<-worker.done
	// }
}

// func bufferedChannel() {
// 	c := make(chan int, 3)
// 	go worker(0, c)
// 	c <- 'a'
// 	c <- 'b'
// 	c <- 'c'
// 	c <- 'd'

// 	time.Sleep(time.Microsecond)
// }

// func channelClose() {
// 	c := make(chan int)
// 	go worker(0, c)
// 	c <- 'a'
// 	c <- 'b'
// 	c <- 'c'
// 	c <- 'd'
// 	close(c)
// 	time.Sleep(time.Microsecond)
// }

func main() {
	chanDemo()
}
