package main

import "fmt"

func main() {
	names := [5]string{"leeda", "dhlee", "DaHyeon"}

	names[3] = "hahaha"
	names[4] = "hahaha"

	fmt.Println(names)

	names2 := []string{"leeda", "dhlee", "DaHyeon"}
	names2 = append(names2, "lalala")

	fmt.Println(names2)
}
