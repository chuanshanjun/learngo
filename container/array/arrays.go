package main

import "fmt"

// []int 切片  [5]int 数组
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	// 1 定义一个指定长度的数组
	var arr1 [5]int
	// 2 定义一个指定长度的数组，并且赋上初始值
	arr2 := [3]int{1, 3, 5}
	// 3 数组长度用...表示 编译器自动识别，后面赋上初始值
	arr3 := [...]int{2, 4, 6, 8, 10}
	// 4 二维数组 4行5列
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// 使用最原始的for来遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	// 用range来遍历
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	// 用range来接受下标和值
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// 只接受值不接收下标
	for _, v := range arr3 {
		fmt.Println(v)
	}

	printArray(arr1)
	// arr2 是长度为3的数组 [3]int 但 printArray()接收的是长度为5的数组
	// 所以认为这两个不是一个类型，无法执行
	//printArray(arr2)
	printArray(arr3)

	// go的数组是值传递,所以在方法里面修改arr[0] 原始的值也不会变
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	fmt.Println("===>>>")

	// printArray2(arr *[5]int)方法参数类型是数组指针
	// 所以传入arr3的地址，收到地址后对该数组进行操作
	// 其实就是对一个数组进行操作，所以修改后会影响原先的数组值
	printArray2(&arr3)
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
}

// 参数类型是 数组指针
func printArray2(arr *[5]int) {
	arr[0] = 100
	for i,v := range arr {
		fmt.Println(i, v)
	}
}
