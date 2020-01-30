package mock

import "fmt"

// 结构的名字就不要叫MockRetriver了，因为包名已经有mock了
// 整个包中都没出现Retriver接口的信息
// 我们只要实现了Get方法，那么我们就认为它实现了这个接口
type Retriver struct {
	Contents string
}

// 这个string接口，相当于java中的ToString
// 实现string接口,实现自己需要的内容
func (r *Retriver) String() string {
	return fmt.Sprintf(
		"Retriver: {Contents=%s}", r.Contents)
}

// 同理Retriver实现接口不需要声明，只要直接声明这个方法就可以
// (r Retriver) 是值接收者，值接收者是不会修改它的值的
func (r *Retriver) Post(url string,
	form map[string]string) string {
	r.Contents = form["contents"] // 将map中的"contents"的值赋予进去
	return "ok"
}

// Retriver实现Retriver接口
func (r *Retriver) Get(url string) string {
	return r.Contents
}
