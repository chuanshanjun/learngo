package main

import (
	"fmt"
	"time"
)

// goroutine 版本1
//func main() {
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			for {
//				fmt.Printf("Hello from"+
//					"goroutine %d\n", i)
//			}
//		}(i)
//	}
//	time.Sleep(time.Millisecond)
//}

//func main() {
//	var a [10]int
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			for {
//				a[i]++
//				runtime.Gosched()
//			}
//		}(i)
//	}
//	time.Sleep(time.Millisecond)
//	fmt.Println(a)
//}

// 不把i传进去就相当于 func就相当于是一个闭包
// 读取的是外面的i，当外面的main跑完后 i变成10
// 然后里面读取i a[10]就超标了
//func main() {
//	var a [10]int
//	for i := 0; i < 10; i++ {
//		go func() {
//			for {
//				a[i]++
//				runtime.Gosched()
//			}
//		}()
//	}
//	time.Sleep(time.Millisecond)
//	fmt.Println()
//}

// 定义1000个协程运行
func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello From"+
					"goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
