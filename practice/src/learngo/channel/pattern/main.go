package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string) <-chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service %s: message %d", name, i)
			i++
		}
	}()
	return c
}

// 接收两个channel的数据,依次输出到新的channel
func fanIn(chs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, ch := range chs {
		go func(in <-chan string) {
			for {
				c <- <-in
			}
		}(ch) // 传入ch的复制,防止go routine冲突
	}
	return c
}

// select实现,仅需一个go routine
func fanInBySelect(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()
	return c
}

// 非阻塞等待
func nonBlockingWait(c <-chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

// 超时等待
func timeoutWait(c <-chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout):
		return "", false
	}
}

// 主动中断任务
func msgGenInterrup(name string, done chan struct{}) <-chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(5000)) * time.Millisecond):
				c <- fmt.Sprintf("service %s: message %d", name, i)
			case <-done:
				fmt.Println("cleaning up")
				time.Sleep(2 * time.Second)
				fmt.Println("cleaning done")
				done <- struct{}{}
				return
			}
			i++
		}
	}()
	return c
}

func main() {
	// m1 := msgGen("s1")
	// m2 := msgGen("s2")
	// m := fanIn(m1, m2)
	// // m := fanInBySelect(m1, m2)
	// for {
	// 	// fmt.Println(<-m1)
	// 	// fmt.Println(<-m2)
	// 	fmt.Println(<-m)
	// }

	/* ============== 并发任务控制 ================== */
	// m1 := msgGen("s1")
	// // m2 := msgGen("s2")
	// // for {
	// // 	fmt.Println(<-m1)
	// // 	if m, ok := nonBlockingWait(m2); ok {
	// // 		fmt.Println(m)
	// // 	} else {
	// // 		fmt.Println("no message from s2")
	// // 	}
	// // }
	// for {
	// 	if m, ok := timeoutWait(m1, 2*time.Second); ok {
	// 		fmt.Println(m)
	// 	} else {
	// 		fmt.Println("timeout")
	// 	}
	// }

	done := make(chan struct{})
	m1 := msgGenInterrup("s1", done)
	for i := 0; i < 5; i++ {
		if m, ok := timeoutWait(m1, time.Second); ok {
			fmt.Println(m)
		} else {
			fmt.Println("timeout")
		}
	}

	done <- struct{}{}
	// time.Sleep(time.Second)
	// 优雅退出
	<-done
}
