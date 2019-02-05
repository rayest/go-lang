package main

import (
	"flag"
	"./go_03" // 引入包
)

// 输入命令行参数
// 运行：go run go_02.go -name="rayest"

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	invokedByGo02(name)
	go_03.Hello()
}
