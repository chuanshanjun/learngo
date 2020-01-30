package main

import "fmt"

// 函数adder() 没有参数，但有返回值，返回值是一个函数=func(int) int
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := adder()
	fmt.Printf("> %T %v\n", a, a)
	for i := 0; i < 10; i++ {
		// 去累加这个i, a这个函数体中存放sum的数据
		fmt.Printf("0 + 1 + ... %d = %d\n", i, a(i))
	}
}
