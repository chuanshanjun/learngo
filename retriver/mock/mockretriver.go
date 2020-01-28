package mock

// 结构的名字就不要叫MockRetriver了，因为包名已经有mock了
// 整个包中都没出现Retriver接口的信息
// 我们只要实现了Get方法，那么我们就认为它实现了这个接口
type Retriver struct {
	Contents string
}

// Retriver实现Retriver接口
func (r Retriver) Get(url string) string {
	return r.Contents
}
