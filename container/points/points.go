package main

import "fmt"

func main() {
	a := 20

	ip := &a // 指针变量的存储地址

	fmt.Printf("a 变量的地址是: %x\n", &a)

	fmt.Printf("ip 变量存储的指针地址:%x\n", ip)

	// 使用指针访问值
	fmt.Printf("*ip 变量的值:%d\n", *ip)

	// 空指针
	var ptr *int // 整型指针 *int

	fmt.Printf("ptr 的值: %x\n", ptr)

	// go程序设计语言 2.3.2 指针
	x := 1
	p := &x         // p是整型指针,指向x
	fmt.Println(*p) // "1"
	*p = 2          // 等于 x = 2
	fmt.Println(x)  // 结果 2
	*p = 3
	fmt.Printf("打印p 是打印 地址值 %s \n", p)
	fmt.Printf("打印*p 是打印 *p的变量值 %d", *p)
}
