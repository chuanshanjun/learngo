package main

import (
	"fmt"
)

func tryRecover() {
	// 这是一个匿名函数需要执行,所以后面加()
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			// %v 相应值的默认格式
			panic(fmt.Sprintf("I dont't know what to do: %v", r))
		}
	}()

	//panic(errors.New("this is an error"))

	// runtime error 也可以用recover
	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	// 或者可以repanic
	panic(123)
}

func main() {
	tryRecover()
}
