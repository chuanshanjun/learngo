package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am IPhone, I can call you")
}

type Duck interface {
	Quack()  // 鸭子叫
	DuckGo() // 鸭子走
}

type Chicken struct {
}

func (c Chicken) IsChicken() bool {
	fmt.Println("我是小鸡")
	return false
}

func (c Chicken) Quack() {
	fmt.Println("嘎嘎")
}

func (c Chicken) DuckGo() {
	fmt.Println("大摇大摆的走")
}

func DoDuck(d Duck) {
	d.Quack()
	d.DuckGo()
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	chicken := Chicken{}
	// 因为小鸡实现了鸭子的!所有!方法，所以小鸡也是鸭子
	DoDuck(chicken)
	// 小鸡直接嘎嘎叫
	chicken.Quack()

	// 一个接口没有定义任何方法，空接口
	// 任何类型都可以满足它，此时可以给它传任意类型的参数
	// Go的interface{}常常会被作为函数的参数传递,可以帮助我们实现范型的效果
	// Go中接口是一组方法的集合
	var nilI interface{} = 1
	fmt.Println(nilI)
}
