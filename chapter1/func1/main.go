package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("koamania")
	fmt.Println(totalLength, upperName)

	totalLength2, _ := lenAndUpper("koamania")
	fmt.Println(totalLength2)

	repeatMe("koamania", "dhlee", "Leeda", "LeeDaHyeon")
}
