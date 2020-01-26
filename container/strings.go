package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课!" // UTF-8 编码 是可变长的 asiic一个字节 中文3个字节
	fmt.Println(s)
	// %s 输出字符串表示 string or []byte
	fmt.Printf("%s\n", []byte(s))

	// %X 将字符串转换为大写的16进制格式
	fmt.Printf("%X\n", s)

	// 在UTF-8 编码模式下 每个中文是3字节
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	// 所以这个i 每次拿的是 这个字符开头的 索引
	// 如0 1 2 3 6 9 像6和9 就是 爱 和 幕 字符开头的索引位置
	// 所以这样遍历的时候就不方便了
	for i, ch := range s { // ch is a rune 一个4字节的整数 因为此时ch 是int32
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	//Unicode 是「字符集」
	//UTF-8 是「编码规则」

	// 输出8个
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	fmt.Println("===>>>")

	bytes := []byte(s)
	for len(bytes) > 0 {
		// asiic 码返回的时候size是1 中文size 是3
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// 我们将string转成 rune[] 数组 或者 rune slice之后
	// 就像其他语言一样转成char就对了
	// 像rune就比较上层,byte就比较底层
	// rune需要4个字节
	// 所以[]rune 是将它转换成4个字节的数组，是重新开辟了内存空间
	// 并且其中还有utf-8转Unicode的过程
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ",i, ch)
	}
	fmt.Println()
}
