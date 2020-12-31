package main

import "fmt"

type errorString string

//TODO Question
// 标准库的 errors 为什么不这样实现，这样也能保证相同的字符串错误不会导致相等
func (e *errorString) Error() string {
	return string(*e)
}

func New(text string) error {
	var res = errorString(text)
	return &res
}

// 结构体比较的时候会比较字段的值，这样定义 error 会导致错误比较出错
type struErr struct {
	s string
}

func (e struErr) Error() string {
	return e.s
}

func NewStruErr(s string) error {
	return struErr{s}
}

func main() {
	fmt.Println("字符串指针变量定义的 error 比较", New("EOF") == New("EOF"))
	fmt.Println("结构体变量定义的 error 比较", NewStruErr("EOF") == NewStruErr("EOF"))
}
