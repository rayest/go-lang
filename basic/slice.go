package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("Init empty: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("After set: ", s)
	fmt.Println("Get s[2]: ", s[2])

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("After append: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("'c' copied from s: ", c)

	l := s[2:5]
	fmt.Println("First sliced: ", l)

	l = s[:5]
	fmt.Println("Second sliced: ", l)

}
