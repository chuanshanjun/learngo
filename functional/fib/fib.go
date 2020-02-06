package fib

// 斐波那契数列生成器函数，返回一个函数
func Fibonacci() func() int {
	// 给初始值
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
