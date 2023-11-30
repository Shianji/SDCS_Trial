package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 只需要实现POST、GET、DELETE GET：从服务器获取资源，例如网页或图像。POST：将数据提交给服务器进行处理，例如表单提交或文件上传。DELETE：从服务器中删除资源。
func post() {
	r, err := http.Post("http://www.xiaomi.com/post", "", nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", content)
}

func get() {
	r, err := http.Get("http://www.baidu.com/get")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", content)
}

func put() {

}

func delete() {

}

func main() {
	// get()
	post()
}
