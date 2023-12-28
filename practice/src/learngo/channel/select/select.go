package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func doWork(id int, c chan int) {
	for n := range c {
		// 运行速度不统一会造成数据丢失
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

// 创建的channel只能发数据
func createWorker(id int) chan int {
	c := make(chan int)
	go doWork(id, c)
	return c
}

func main() {
	// 默认阻塞
	// var c1, c2 chan int // c1 and c2 = nil
	var c1, c2 = generator(), generator()
	w := createWorker(0)

	n := 0
	// hasValue := false

	var values []int
	tm := time.After(10 * time.Second) // 计时器

	for {
		var activateWorker chan<- int
		var activateValue int
		if len(values) > 0 {
			activateWorker = w
			activateValue = values[0]
		}
		// 非阻塞式
		select {
		case n = <-c1:
			// fmt.Println("Received from c1:", n)
			// hasValue = true
			values = append(values, n)
		case n = <-c2:
			// hasValue = true
			// fmt.Println("Received from c2:", n)
			// default:
			// 	fmt.Println("No value received")
			values = append(values, n)
		case activateWorker <- activateValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tm:
			fmt.Println("GoodBye")
			return
		}
	}
}
