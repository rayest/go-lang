package main

import (
	"fmt"
)

// go run go_02.go go_03.go --name=lee
func invokedByGo02(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
