package main

import "fmt"

func main() {

	aMap := map[string]int{"one": 1, "two": 2, "three": 3}
	k := "two"
	value, ok := aMap[k]
	if ok {
		fmt.Println("value: ", value)
	} else {
		fmt.Println("Not found!")
	}
}
