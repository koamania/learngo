package main

import "fmt"

func main() {
	//var leeda map[string]string = map[string]string{"name": "leeda", "age": "12"}
	leeda := map[string]string{"name": "leeda", "age": "12"}

	fmt.Println(leeda)

	for key, value := range leeda {
		fmt.Println(key, value)
	}


}
