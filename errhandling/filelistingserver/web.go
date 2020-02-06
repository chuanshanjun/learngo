package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chuanshan/learngo/errhandling/filelistingserver/filelisting"
)

// 1 第一版
// 功能是ok，但出错处理做的不好，如path为 /list/fib.txta 报错就不友好
//func main() {
//	http.HandleFunc("/list/",
//		func(writer http.ResponseWriter,
//			request *http.Request) {
//			path := request.URL.Path[len("/list/"):]
//			file, err := os.Open(path)
//			if err != nil {
//				panic(err)
//			}
//			defer file.Close() // 打开后要记得及时关闭
//
//			all, err := ioutil.ReadAll(file)
//			if err != nil {
//				panic(err)
//			}
//
//			writer.Write(all)
//		})
//
//	err := http.ListenAndServe(":8888", nil)
//	if err != nil {
//		panic(err)
//	}
//}

// 第二版
// 对err进行相应捕获，但报错信息是内部的报错信息，需要包装异常
//func main() {
//	http.HandleFunc("/list/",
//		func(writer http.ResponseWriter,
//			request *http.Request) {
//			path := request.URL.Path[len("/list/"):]
//			file, err := os.Open(path)
//			if err != nil {
//				http.Error(writer,
//					err.Error(),
//					http.StatusInternalServerError)
//				// 处理完err后，可直接return
//				return
//			}
//			defer file.Close() // 打开后要记得及时关闭
//
//			all, err := ioutil.ReadAll(file)
//			if err != nil {
//				panic(err)
//			}
//
//			writer.Write(all)
//		})
//
//	err := http.ListenAndServe(":8888", nil)
//	if err != nil {
//		panic(err)
//	}
//}

// 定义 appHandler函数并返回error
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 接收appHandler让其做统一的业务处理，然后再对它返回的err做统一的处理，返回http.HandleFunc所需要的参数
// errWrapper使用的是函数式编程 输入是一个函数 输出也是一个函数
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		// 接收业务逻辑handler返回的err
		err := handler(writer, request)
		// 对err进行分类处理
		if err != nil {
			// 此log不是go语言标准库中的，是gopm中的
			//log.Warn("Error handling request: %s",
			//	err.Error())

			// 处理我们自定义的error
			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}

			// 标准库中的log没有warn这个级别的
			log.Printf("Error handling request: %s",
				err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code), code)
		}
	}
}

type userError interface {
	// error 给系统看的
	error
	// 给用户看的
	Message() string
}

// 第三版
// 拆分业务逻辑
func main() {
	// 我们将业务逻辑放到filelisting.HandlerFileList函数中
	// HandlerFileList函数只做业务逻辑，碰到错误就return
	http.HandleFunc("/",
		errWrapper(filelisting.HandlerFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
