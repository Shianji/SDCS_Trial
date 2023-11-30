package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func test1() {
	p := Person{
		"tom",
		26,
		"tom@qq.com",
	}
	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b))
}

func test2() {
	b := []byte(`{"Name":"tom","Age":26,"Email":"tom@qq.com"}`)
	var p Person
	// var p interface{}
	json.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p)
}

// 还可以用于解析嵌套类型
func test3() {
	b := []byte(`{"Name":"tom","Age":26,"Email":"tom@qq.com","Parent":["big tom1","big tom2"]}`)
	// var f interface{}
	var f map[string]interface{} //这样定义就可以在下面遍历f中的内容了,interface{} 是一种空接口类型,它表示一个空集合。interface{} 不包含任何方法，因此它可以表示任何类型的值。
	json.Unmarshal(b, &f)        //注意这里传入的是f的地址
	fmt.Printf("f: %v\n", f)

	for k, v := range f {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}
}

// 扩展到websocket的应用场景
// 从文件解析json对象
func test4() {
	f, _ := os.Open("a.json")
	defer f.Close()
	d := json.NewDecoder(f)
	var v map[string]interface{}
	d.Decode(&v)
	fmt.Printf("v: %v\n", v)
}

// 将json对象写入文件
func test5() {
	p := Person{
		"tom",
		26,
		"tom@qq.com",
	}
	f, _ := os.OpenFile("a.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	e := json.NewEncoder(f)
	e.Encode(p)
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
}
