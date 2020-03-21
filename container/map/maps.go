package main

import (
	"fmt"
)

type User struct {
	Name   string            `json:"name"`
	Id     string            `json:"id"`
	Assets map[string]string `json:"assets"`
}

func main() {
	// 1 直接定义map 的时候赋予初值
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
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

	fmt.Println("......")

	user := new(User)
	user.Assets = make(map[string]string)
	user.Assets["first"] = "java"
	user.Assets["second"] = "ruby"
	user.Assets["third"] = "go"
	user.Assets["forth"] = "javascript"
	for k, v := range user.Assets {
		fmt.Printf("k is %s, v is %s \n", k, v)
	}

	user2 := &User{
		Name: "cool",
		Id:   "1",
		Assets: map[string]string{
			"first":  "java",
			"second": "javascript",
			"third":  "go",
		},
	}

	fmt.Println(user2)

	// 创建一个新的结构体
	fmt.Println(Books{"go", "www.runoob.com", "go teach", 6495407})

	// 使用key -> value 格式
	fmt.Println(Books{
		title:   "go language",
		author:  "www.ruboob.com",
		subject: "go 教程",
		bookId:  6495407,
	})

	// 申明Book1
	var book1 Books
	// 申明Book2
	var book2 Books

	book1.title = "Go 语言"
	book1.author = "www.runoob.com"
	book1.subject = "go 语言教程"
	book1.bookId = 21344

	book2.title = "python 教程"
	book2.author = "www.runoob.com"
	book2.subject = "Python 教程"
	book2.bookId = 88888

	fmt.Printf("book1 is %s", book1)
	fmt.Printf("book1 is %s", book2)

}

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}
