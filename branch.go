package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		// %s 字符串或者byte[]数组的占位符
		fmt.Printf("%s\n", contents)
	}

	// go语言可以将if变成for类的写法，但后面要跟上分号
	if contents2, err2 := ioutil.ReadFile(filename); err2 != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents2)
	}

	fmt.Println(
		grade(0),
		grade(1),
		grade(78),
		grade(100),
	//grade(120),
	)
}

// switch会自动break, 除非使用fallthrough
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		// Sprintf返回一个格式化的字符串但不输出
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}
