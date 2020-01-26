package main

import "fmt"

func main() {
	// 1 直接定义map 的时候赋予初值
	m := map[string]string {
		"name": "ccmouse",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}

	// 2 使用make来定义
	m2 := make(map[string]int) // m2 == empty map

	// 3 使用var来定义
	var m3 map[string]int // m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	// 打印k v
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 只打印k
	for k := range m {
		fmt.Println(k)
	}

	// 只打印v
	for _, v := range m {
		fmt.Println(v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	causeName := m["cause"]
	// key 不存在的话给一个zero value
	fmt.Println(causeName)

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key dose not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
