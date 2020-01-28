package main

import (
	"fmt"
	"time"

	"github.com/chuanshan/learngo/retriver/mock"
	"github.com/chuanshan/learngo/retriver/real"
)

// 接口类型
type Retriver interface {
	Get(url string) string
}

// download是我们的使用者
func download(r Retriver) string {
	return r.Get("http://www.imooc.com")
}

func inspect(r Retriver) {
	fmt.Printf("%T %v\n", r, r)
	fmt.Println("Type Switch:")
	switch v := r.(type) {
	case mock.Retriver:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriver:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

func main() {
	var r Retriver
	// 值接收者，可以接收值
	r = mock.Retriver{"this is a fake www.imooc.com"}
	// 值接收者，也可以接收地址
	//r = &mock.Retriver{"this is a fake www.imooc.com"}
	// %T 相应值的类型的Go语法表示
	// %v 相应值的默认格式
	fmt.Printf("%T %v\n", r, r)
	inspect(r)
	//fmt.Println(download(r))

	mockRetriver := r.(mock.Retriver)
	fmt.Println(mockRetriver.Contents)

	// 也可以直接这么打印
	//fmt.Println(download(mock.Retriver{"this is a fake www.imooc.com2"}))
	// 因为real.Retriver中的Get方法的接收者是指针，所以此处我们赋地址值
	r = &real.Retriver{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))

	inspect(r)

	// TypeAssert 通过.(类型名字) 来获取接口中具体类型
	if realRetriver, ok := r.(*real.Retriver); ok {
		fmt.Println(realRetriver.TimeOut)
	} else {
		fmt.Println("not a mock retriever")
	}

}
