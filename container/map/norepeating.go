package main

import "fmt"

// 不含重复的子字符串的国际版
func lengthOfNonRepeatingSubStrByRune(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	// 用range 对字符串 s 操作 ch 变成了int
	// 使用了[]byte操作后 ch 就变成了byte
	for i, ch := range []rune(s) {
		// map中取不到key的时候会返回zero value
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	// 用range 对字符串 s 操作 ch 变成了int
	// 使用了[]byte操作后 ch 就变成了byte
	for i, ch := range []byte(s) {
		// map中取不到key的时候会返回zero value
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	// asiic编码用一个字节
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	// 中文 UTF-8 编码 占据3个字节
	fmt.Println(lengthOfNonRepeatingSubStrByRune("我爱慕课网我爱网"))
}
