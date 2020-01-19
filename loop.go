package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	// 首先打开文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	// 没有开始条件 没有终止条件
	// 相当与while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 相当于无限循环
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	//forever()
	printFile("abc.txt")
	fmt.Println(
		convertToBin(10), // 1010
		convertToBin(13), // 1101
		convertToBin(23123),
		convertToBin(0),
	)
}
