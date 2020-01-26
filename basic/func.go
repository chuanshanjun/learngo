package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// go 参数在前 类型在后， 返回值也在后
// 同类的参数可以用 ,逗号 区分
// 多返回值一般不要乱用，一般返回一个值还有一个err
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		//return a / b
		// div返回两个返回值，go语言中如果一个变量不需要的话使用_下划线
		q, _ := div(a, b) // q, r := div(a, b)
		return q
	default:
		panic("unsupported operation: " + op)
	}
}

// 用panic的话整个程序中断比较难看，所以我们使用error
func eval2(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

// go语言可以有两个返回值
func div(a, b int) (int, int) {
	return a / b, a % b
}

// 返回值也可以给变量名，这样子接返回值也好接，可以用IDE自动生成去接
// 推荐返回方式 用这种
func div2(a, b int) (q, r int) {
	return a / b, a % b
}

// 这种返回方式也可以，但是函数值一长就搞不清楚了
func div3(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

// 使用函数式编程，把函数做为参数传入
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()    // 反射拿到这个这个函数的指针
	opName := runtime.FuncForPC(p).Name() // runtime 获取函数名
	fmt.Printf("Calling function %s with args"+
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// go语言没有方法重载之类的，只有变长参数
func sum(numbers ...int) int {
	s := 0
	// range 遍历数组中每个数
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) {
	b, a = a, b
}

// 传入指针类型
func swap2(a, b *int) {
	*b, *a = *a, *b
}

func swap3(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	fmt.Println(div(4, 3))
	q, r := div2(3, 7) // 用这种方式接受不了两个返回值
	fmt.Println(q, r)
	fmt.Println(div3(3, 4))
	fmt.Println(eval(3, 4, "/"))
	fmt.Println(eval2(3, 4, "X"))

	// 外面可以来判断error的方法
	if result, err := eval2(3, 5, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// 输出 	第一个main是包名， 第二个pow是函数名
	fmt.Println(apply(pow, 3, 3))

	// 这样就写成匿名函数了
	//输出 Calling function main.main.func1 with args(3, 4)
	// 第一个main是包名， 第二个main.func1 是函数名 因为是匿名函数 所以不知道叫啥 就叫func1
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	swap(a, b)
	// 此时是值传递,原来的a,b值不变,无法交换
	fmt.Println(a, b)

	// 此时要传递指针类型，所以用&表示取地址值
	swap2(&a, &b)
	fmt.Println(a, b)

	c, d := 3, 4
	fmt.Println(swap3(c, d))
}
