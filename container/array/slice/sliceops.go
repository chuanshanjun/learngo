package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n",
		s, len(s), cap(s))
}

func sliceOps() {
	fmt.Println("Creating slice")
	var s []int //Zero value for slice is nil
	// cap每次装不下后 原先容量*2
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	// 制造一个容量和len都是16的slice
	s2 := make([]int, 16)
	// 制造一个len是10，容量是16的slice
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Printf("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Printf("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)
}
