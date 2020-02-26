package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 201212

	fmt.Println(**&b)
}
