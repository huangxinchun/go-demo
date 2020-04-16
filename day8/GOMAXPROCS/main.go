package main

import (
	"fmt"
	"time"
	"runtime"
)

//Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）。

//Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。

//Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。

func a() {
	for i := 1; i < 10; i++{
		go fmt.Printf("a : %v\n", i)
	}
}

func b() {
	for i := 1; i < 10; i++{
		go fmt.Printf("b : %v\n", i)
	}
}
func main() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
