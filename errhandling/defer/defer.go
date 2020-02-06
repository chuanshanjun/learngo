package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/chuanshan/learngo/functional/fib"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("err")
}

func tryDefer2() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 10 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	// 写文件之前要先打开文件
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	// 文件写完后 需要将它关闭
	defer file.Close()

	// 写文件直接写会比较慢，所以直接使用buf.io
	// 这样写是带着buf的先写到内存，然后满了后一下子写出去
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func writeFile2(filename string) {
	// 写文件之前要先打开文件
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	// 自定义error
	err = errors.New("this is a custom error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	// 文件写完后 需要将它关闭
	defer file.Close()

	// 写文件直接写会比较慢，所以直接使用buf.io
	// 这样写是带着buf的先写到内存，然后满了后一下子写出去
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//writeFile("fib.txt")
	//tryDefer2()

	writeFile2("fib.txt")
}
