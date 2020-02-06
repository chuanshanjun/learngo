package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// 第二个版本
//func HandlerFileList(writer http.ResponseWriter,
//	request *http.Request) {
//	path := request.URL.Path[len("/list/"):]
//	file, err := os.Open(path)
//	if err != nil {
//		http.Error(writer,
//			err.Error(),
//			http.StatusInternalServerError)
//		// 处理完err后，可直接return
//		return
//	}
//	defer file.Close() // 打开后要记得及时关闭
//
//	all, err := ioutil.ReadAll(file)
//	if err != nil {
//		panic(err)
//	}
//
//	writer.Write(all)
//}

const prefix = "/list/"

type userError string

// userError接口定义在web.go
func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

// 第三个版本 函数返回一个error
func HandlerFileList(writer http.ResponseWriter,
	request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		// 返回自定义异常
		return userError("path must start with" + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		// 有错误就直接返回
		return err
	}
	defer file.Close() // 打开后要记得及时关闭

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	// 没错就直接返回一个nil
	return nil
}
