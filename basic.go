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

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	triangle()
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
