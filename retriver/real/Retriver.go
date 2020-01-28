package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriver struct {
	UserAgent string
	TimeOut   time.Duration
}

// 将接收者改为指针，通过指针来访问这个get
// 当我们这个struct很大，不想通过普通的传值去拷贝
// 所以我们使用指针
func (r *Retriver) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	// 要记得关闭resp
	resp.Body.Close()
	if err != nil {
		panic(err)
	}

	return string(result)
}
