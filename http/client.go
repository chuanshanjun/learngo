package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// case3 自定义client，查看重定向内容
func main() {
	request, err := http.NewRequest(
		http.MethodGet,
		"http://www.imooc.com", nil)
	request.Header.Add("User-Agent", ""+
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	// http.Get方法要求主动关闭resp.Body
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}

// case2 使用DefaultClient来处理
//func main() {
//	request, err := http.NewRequest(
//		http.MethodGet,
//		"http://www.imooc.com", nil)
//	request.Header.Add("User-Agent", ""+
//		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
//	resp, err := http.DefaultClient.Do(request)
//	if err != nil {
//		panic(err)
//	}
//
//	// http.Get方法要求主动关闭resp.Body
//	defer resp.Body.Close()
//
//	s, err := httputil.DumpResponse(resp, true)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Printf("%s\n", s)
//}

// case1 直接用http.Get请求
//func main() {
//	resp, err := http.Get("HTTP://www.imooc.com")
//	if err != nil {
//		panic(err)
//	}
//
//	// http.Get方法要求主动关闭resp.Body
//	defer resp.Body.Close()
//
//	s, err := httputil.DumpResponse(resp, true)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Printf("%s\n", s)
//}
