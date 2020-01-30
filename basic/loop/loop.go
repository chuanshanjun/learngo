package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

	printFileContents(file)

	//scanner := bufio.NewScanner(file)
	//
	//// 没有开始条件 没有终止条件
	//// 相当与while
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
}

func printFileContents(read io.Reader) {
	scanner := bufio.NewScanner(read)

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

	// ``是跨行字符串
	s := `abc"d"
	lllkkk
	
	123
	k`

	// 打印文件的地方用 read后 适用范围就更加广了
	// btyes.NewReader()
	printFileContents(strings.NewReader(s))
}
