package main

import (
	"fmt"
	"flag"
)

// 输入命令行参数
// 运行：go run go_02.go -name="rayest"

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()

	fmt.Printf("Hello, %s!\n", name)
}
