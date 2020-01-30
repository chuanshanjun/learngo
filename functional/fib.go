package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 斐波那契数列生成器函数，返回一个函数
func fibonacci() intGen {
	// 给初始值
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 在go语言中是个type就可以实现接口
// go语言呢的特色，给函数也可以实现接口
type intGen func() int

// intGen这个类型,实现了Read接口
func (g intGen) Read(p []byte) (n int, err error) {
	// 我们去调用一下函数就可以取得它的元素
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	// 将这个数转为字符串,然后将s写进这个byte[]中
	s := fmt.Sprintf("%d\n", next)
	// strings.NewReader(s).Read(p) 来代理Read()
	return strings.NewReader(s).Read(p)
}

func printFileContents(read io.Reader) {
	scanner := bufio.NewScanner(read)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	// printFileContents() 入参是一个Reader接口类型
	// 此时我们的f也就是斐波那契数列生成器实现了Reader接口
	printFileContents(f)
}
