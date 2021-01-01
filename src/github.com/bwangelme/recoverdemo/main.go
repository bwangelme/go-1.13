// 校验 goroutine 的 panic 主进程的 recover 是抓不到的

package main

import (
	"fmt"
	"time"
)

func main() {
	//f()
	//fmt.Println("Returned normally from f.")

	NewRoutinePanic()
	fmt.Println("Returned normally from NewRoutinePanic.")
}

// Revoer 仅仅定义在 defer 函数中才会生效，此时 revoer 会获取到panic 传递的值，并将程序的控制流程拉回到正常流程上来。
// 如果在普通函数中执行 revoer，因为 panic 还没有发生，此时执行 recover 不会收到任何值。
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

// 如果 panic 和 recover 不在同一个 goroutine 中，那么 recover 是捕获不到 panic 的异常的
func NewRoutinePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in NewRoutinePanic", r)
		}
	}()

	fmt.Println("Start Calling g")
	go g(4)
	time.Sleep(1 * time.Second)
	fmt.Println("End Calling g")
}

// Panic 发生的时候，panic 之后的流程都不会执行了，但是函数注册的 defer 函数仍然会执行
func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
    fmt.Println("End of calling g", i)
}
