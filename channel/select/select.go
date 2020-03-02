package main

import (
	"fmt"
	"math/rand"
	"time"
)

// case2
// for 循环select 哪个有值收哪个
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

func main() {
	// 默认c1 c2是没有值的，那么输出default
	// 默认channel发数据或者是收数据都会阻塞住
	// 如果想非阻塞就是select加default
	var c1, c2 = generator(), generator()
	for {
		select {
		case n := <-c1:
			fmt.Println("Received from c1:", n)
		case n := <-c2:
			fmt.Println("Received from c2:", n)
		}
	}
}

// case1
//func main() {
//	// 默认c1 c2是没有值的，那么输出default
//	// 默认channel发数据或者是收数据都会阻塞住
//	// 如果想非阻塞就是select加default
//	var c1, c2 chan int // c1 and c2 = nil
//	for {
//		select {
//		case n := <-c1:
//			fmt.Println("Received from c1:", n)
//		case n := <-c2:
//			fmt.Println("Received from c2:", n)
//		default:
//			fmt.Println("No value received")
//		}
//	}
//}
