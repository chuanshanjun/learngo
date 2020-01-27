package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 包内变量只能使用var来赋值
var aa = 3

// 也可使用 var () 批量赋值
var (
	bb = 4
	cc = 5
	dd = 6
)

// 包内变量不能使用:= 赋值
//bb := 4

// 常量同样可以定义为包变量
const filename = "abc.txt"

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	triangle()
	consts()
	fmt.Println(filename)
	enums()
	enums2()
	enums3()
	enums4()
}

func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)
	// 单引号围绕的字符字面值，由Go语法安全地转义
	fmt.Printf("%d %q \n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	// 在方法内推荐使用此种类型的赋值方式
	a, b, c, s := 3, 4, false, "def"
	// 上面用 := 赋值过的变量后，下面不能再用:= 赋值了
	//a := 4
	fmt.Println(a, b, c, s)
}

func euler() {
	// 有小数点而无指数，例如 123.456
	// cmplx 包为复数提供基本的常量和数学函数
	fmt.Printf("%.3f\n", cmplx.Exp((1i*math.Pi)+1))
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	// 常量你可以不指定类型，他相当于一个文本替换
	const filename = "abc.txt"
	const a, b = 3, 4

	// 同样常量也可以用 () 包裹
	// go语言中的常量一般不大写
	const (
		filename2 = "ccc"
		e, f      = 3, 5
	)
	var c int
	// 此处常量自己转换成float
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c, filename)
}

// go语言中没有枚举关键字一般用常量
func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)
}

// iota为自增可以从0开始自增
func enums2() {
	const (
		cpp    = iota
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)
}

func enums3() {
	const (
		cpp    = iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp, javascript, python, golang)
}

// iota可以做为自增值的表达式的种子
func enums4() {
	const (
		b    = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, pb)
}
