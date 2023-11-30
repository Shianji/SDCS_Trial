package main

import (
	"errors"
	"fmt"
	"time"
)

func check(s string) (string, error) {
	if s == "" {
		err := errors.New("字符串不能为空") //errors.New返回的是error类型，是一个接口类型，但实际实现中，返回的是一个实现了error接口中Error()string方法的结构体
		return "", err
	} else {
		return s, nil
	}
}

// 自定义错误
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(2023, 11, 14, 20, 45, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func main() {
	s, err := check("hello world")
	if err != nil {
		fmt.Printf("err: %v\n", err) //打印err时，两种输出方式都可以
		fmt.Printf("err: %v\n", err.Error())
	} else {
		fmt.Printf("s: %v\n", s)
	}
	fmt.Println("--------------------------------")
	if err = oops(); err != nil {
		fmt.Println(err) //当打印err时，实际上是调用了MyError类型的Error()方法，将其转换为字符串并进行打印。
	}
}
