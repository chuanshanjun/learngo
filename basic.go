package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
}

func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)
	fmt.Printf("%d %q", a, s);
}
