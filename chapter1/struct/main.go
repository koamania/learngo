package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string
}

func main() {

	// map
	//leeda := map[string]string{"name": "leeda"}
	favFood := []string{"kimchi", "bulgogi"}
	leeda := person {"leeda", 32, favFood}

	fmt.Println(leeda)

	leeda2 := person{name:"leeda", age:32, favFood: []string{"kimchijjigae"}}

	fmt.Println(leeda2)
}
