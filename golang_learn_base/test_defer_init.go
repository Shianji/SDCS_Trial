package main

import "fmt"

func f() {
	fmt.Println("start.......")
	fmt.Println("step1.......")
	fmt.Println("step2.......")
	fmt.Println("step3.......")
	fmt.Println("end.......")
}

// defer会将起后面跟随的语句延迟处理，在defer归属的函数即将返回前，将defer延迟处理的语句按照逆序执行，想当于出入栈
// defer可以用来关闭文件句柄，锁资源释放等，例如在打开文件后使用defer关闭文件会等到程序结束再执行关闭操作
func f_defer() {
	fmt.Println("start.......")
	defer fmt.Println("step1.......")
	defer fmt.Println("step2.......")
	defer fmt.Println("step3.......")
	fmt.Println("end.......")
}

var i int = initVar()

func init() {
	fmt.Println("init......")
}

func init() {
	fmt.Println("init2......")
}

func initVar() int {
	fmt.Println("initVar......")
	return 100
}

// go语言程序初始化顺序：变量初始化->init()->main()
// init优先于main函数自动执行，init不能被其他函数调用，它没有输入参数和返回值，每个包和每个源文件都可以有多个init函数
func main() {
	f()
	fmt.Println("-------------------")
	f_defer()
}
