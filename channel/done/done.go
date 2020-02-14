package main

import (
	"fmt"
	"sync"
)

//func doWork(id int, in chan int, done chan bool) {
//	for n := range in {
//		fmt.Printf("worker %d received %c\n",
//			id, n)
//		// 同样我在这发done，必须外面也要收done
//		go func() {
//			done <- true
//		}()
//	}
//}
//
//type worker struct {
//	in   chan int
//	done chan bool
//}
////
//func createWorker(id int) worker {
//	w := worker{
//		in:   make(chan int),
//		done: make(chan bool),
//	}
//	go doWork(id, w.in, w.done)
//	return w
//}
//
//func chanDemo() {
//	var workers [10]worker
//	for i := 0; i < 10; i++ {
//		workers[i] = createWorker(i)
//	}
//	for i, worker := range workers {
//		// 所有的channel发 都是阻塞式的，另一端一定要有人收
//		worker.in <- 'a' + i
//	}
//	for i, worker := range workers {
//		worker.in <- 'A' + i
//	}
//	// wait for all of them
//	for _, worker := range workers {
//		// 因为上面发了两次，所以要收两次
//		<-worker.done
//		<-worker.done
//	}
//}

// 我们除了将doWork中在用一个go func() {done <- true}()
// 这种方法外，我们还可以用下面的分开等
//func doWork(id int, in chan int, done chan bool) {
//	for n := range in {
//		fmt.Printf("worker %d received %c\n",
//			id, n)
//		done <- true // 同样我在这发done，必须外面也要收done
//	}
//}
//
//type worker struct {
//	in   chan int
//	done chan bool
//}
//
//func createWorker(id int) worker {
//	w := worker{
//		in:   make(chan int),
//		done: make(chan bool),
//	}
//	go doWork(id, w.in, w.done)
//	return w
//}
//func chanDemo() {
//	var workers [10]worker
//	for i := 0; i < 10; i++ {
//		workers[i] = createWorker(i)
//	}
//	for i, worker := range workers {
//		worker.in <- 'a' + i // 所有的channel发 都是阻塞式的，另一端一定要有人收
//	}
//	// wait for all of them
//	for _, worker := range workers {
//		// 因为上面发了两次，所以要收两次
//		// 这个就是分开等，小写字母都发完后，我再等小写的全发完
//		<-worker.done
//	}
//	for i, worker := range workers {
//		worker.in <- 'A' + i
//	}
//	// wait for all of them
//	for _, worker := range workers {
//		// 因为上面发了两次，所以要收两次
//		// 这个等大写的字母
//		<-worker.done
//	}
//}

// 或者是使用wg *sync.WaitGroup
// 我们知道要20个任务，所以增加20个任务，没完成一个任务就done
// 这样到最后wg.wait()满20个它就退出了
// 要注意用的是一个wg所以是引用传递，传的是指针
//func doWork(id int, in chan int, wg *sync.WaitGroup) {
//	for n := range in {
//		fmt.Printf("worker %d received %c\n",
//			id, n)
//		// 同样我在这发done，必须外面也要收done
//		wg.Done()
//	}
//}
//
//type worker struct {
//	in chan int
//	wg *sync.WaitGroup
//}
//
//// waiGroup是只有一个的，大家共用这一个group
//func createWorker(
//	id int, wg *sync.WaitGroup) worker {
//	w := worker{
//		in: make(chan int),
//		wg: wg,
//	}
//	go doWork(id, w.in, wg)
//	return w
//}
//
//func chanDemo() {
//	var workers [10]worker
//	var wg sync.WaitGroup
//	for i := 0; i < 10; i++ {
//		workers[i] = createWorker(i, &wg)
//	}
//
//	wg.Add(20)
//
//	for i, worker := range workers {
//		// 所有的channel发 都是阻塞式的，另一端一定要有人收
//		worker.in <- 'a' + i
//	}
//	for i, worker := range workers {
//		worker.in <- 'A' + i
//	}
//	wg.Wait()
//}

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n",
			id, n)
		// 同样我在这发done，必须外面也要收done
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

// waiGroup是只有一个的，大家共用这一个group
func createWorker(
	id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers {
		// 所有的channel发 都是阻塞式的，另一端一定要有人收
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	chanDemo()
}
