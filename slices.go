package main

import "fmt"

func updateSlice(arr []int) {
	arr[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]

	// slice本身没有数据，是对底层array的一个view
	fmt.Println("arr[2:6]", s)
	fmt.Println("arr[:6]", arr[:6])
	fmt.Println("arr[2:]", arr[2:])
	fmt.Println("arr[:]", arr[:])

	s1 := arr[2:]
	fmt.Println("s1 = ", s1)
	s2 := arr[:]
	fmt.Println("s2 = ", s2)

	// 所以对slice操作会影响底层的array
	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println("s1", s1)
	fmt.Println("arr", arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println("s2", s2)
	fmt.Println("arr", arr)

	fmt.Println("Reslice")
	fmt.Println("s2", s2)
	s2 = s2[:5]
	fmt.Println("s2", s2)
	s2 = s2[2:]
	fmt.Println("s2", s2)

	// 用[:6]这种方法slice会自动向后"扩容"，把后面的东西打印出来
	// 这种打印方式是用的cap
	fmt.Println("s2[:6]", s2[:6])
	//fmt.Println("s2[:6]", s2[:7])
	// 用[5]打印，使用的是len
	//fmt.Println("s2[6]", s2[5])

	// %v 相应值的默认格式 %d 十进制表示
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	s6 := append(s5, 13)
	fmt.Println("s3, s4, s5, s6", s3, s4, s5, s6)
	// 此时arr 只能输出到12 s6新增的13原来的arr中没有
	// 此时s6 view的不是原来的arr而是一个新的arr
	// go会开辟一个新的arr 然后将原来的arr拷贝过去
	// 所以我们需要一个新的slice来接这个
	fmt.Println("arr =", arr)
}
