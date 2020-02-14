package main

import (
	"fmt"
	"time"
)

// 此时是第一版～此时里面的go func 相当与一个闭包，因为引用了外面的c
//func chanDemo() {
//	// 使用make制作一个chan 里面装载的是int类型
//	c := make(chan int)
//	go func() {
//		for {
//			n := <-c
//			fmt.Println(n)
//		}
//	}()
//	c <- 1
//	c <- 2
//	// 休眠会不然可能收完2后，来不及打印随着main的goroutine的退出，其他的goroutine也会退出
//	time.Sleep(time.Millisecond)
//}

// 第二版 将匿名函数提取出来
//func worker(c chan int) {
//	for {
//		n := <-c
//		fmt.Println(n)
//	}
//}
//
//func chanDemo() {
//	c := make(chan int)
//	go worker(c)
//	c <- 1
//	c <- 2
//	time.Sleep(time.Millisecond)
//}

// 版本2.1
//func worker(id int, c chan int) {
//	for {
//		// 不需要中间变量n,直接使用<-c
//		// n := <-c
//		// %c 相应Unicode码点所表示的字符
//		fmt.Printf("worker %d receieved %c\n",
//			id, <-c)
//	}
//}
//
//func chanDemo() {
//	var channels [10]chan int
//	for i := 0; i < 10; i++ {
//		channels[i] = make(chan int)
//		go worker(i, channels[i])
//	}
//
//	for i := 0; i < 10; i++ {
//		channels[i] <- 'a' + i
//	}
//
//	// 每个worker都在不停的收，所以我们还可以继续发送
//	// 如上面打了a没和下面的大A连续，是因为中间有print
//	// print是IO操作，中间会有切换
//	for i := 0; i < 10; i++ {
//		channels[i] <- 'A' + i
//	}
//	time.Sleep(time.Millisecond)
//}

// 版本2.2 worker->createWorker
// 此函数负责创建一个channel，并把要做的逻辑写好
// 然后将这个channel返回出去
//func createWorker(id int) chan int {
//	c := make(chan int)
//	go func() {
//		for {
//			fmt.Printf("worker %d received %c\n",
//				id, <-c)
//		}
//	}()
//	return c
//}
//
//func chanDemo() {
//	var channels [10]chan int
//	for i := 0; i < 10; i++ {
//		channels[i] = createWorker(i)
//	}
//
//	for i := 0; i < 10; i++ {
//		channels[i] <- 'a' + i
//	}
//
//	for i := 0; i < 10; i++ {
//		channels[i] <- 'A' + i
//	}
//}
//
//func main() {
//	chanDemo()
//}

// 2.3
// channel 是用来接收数据的 所以 chan<-
//func createWorker(id int) chan<- int {
//	c := make(chan int)
//	go func() {
//		for {
//			fmt.Printf("worker %d received %c\n",
//				id, <-c)
//		}
//	}()
//	return c
//}
//
//func chanDemo() {
//	var channels [10]chan<- int
//	for i := 0; i < 10; i++ {
//		channels[i] = createWorker(i)
//	}
//
//	for i := 0; i < 10; i++ {
//		channels[i] <- 'a' + i
//		// 此时我就不可以从channels中收数据了
//		// 提示是说 此时channel 是send only
//		//n := <-channels[i]
//	}
//
//	for i := 0; i < 10; i++ {
//		channels[i] <- 'A' + i
//	}
//}
//
//func worker(id int, c chan int) {
//	// <-c channel 发送数据出去
//	for {
//		fmt.Printf("Worker %d received %c\n",
//			id, <-c)
//	}
//}
//
//func bufferedChannel() {
//	// c的buffered缓冲是3，所以先给它三个数据
//	c := make(chan int, 3)
//	go worker(0, c)
//	c <- 'a'
//	c <- 'b'
//	c <- 'c'
//	time.Sleep(time.Millisecond)
//}
//
//func closeWorker(id int, c chan int) {
//	for {
//		n, ok := <-c
//		if !ok {
//			break
//		}
//		fmt.Printf("Worker %d received %c\n",
//			id, n)
//	}
//}

// 使用range，当c发完的时候，程序就跳出来啊
//func closeWorker2(id int, c chan int) {
//	for n := range c {
//		fmt.Printf("Worker %d received %c\n",
//			id, n)
//	}
//}
//
//func channelClose() {
//	c := make(chan int, 3)
//	go closeWorker(0, c)
//	c <- 'a'
//	c <- 'b'
//	c <- 'c'
//	// channel close后channel还收可以接收数据的
//	// 接收数据是int的初始值0 string就是空串
//	close(c)
//	time.Sleep(time.Millisecond)
//}

// 使用range的时候c里面没有数据就结束循环
//func closeWorker2(id int, c chan int) {
//	for n := range c {
//		fmt.Printf("Worker %d received %c\n",
//			id, n)
//	}
//}
//
//func channelClose() {
//	c := make(chan int, 3)
//	go closeWorker2(0, c)
//	c <- 'a'
//	c <- 'b'
//	c <- 'c'
//	// channel close后channel还收可以接收数据的
//	// 接收数据是int的初始值0 string就是空串
//	close(c)
//	time.Sleep(time.Millisecond)
//}

func work(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d received %c\n",
			id, n)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go work(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
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

func main() {
	//chanDemo()
	//bufferedChannel()
	//channelClose()
	chanDemo()
}
