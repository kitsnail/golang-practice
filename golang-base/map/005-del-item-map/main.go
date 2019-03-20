package main

import "fmt"

func main() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println("Deleting values")
	name, ok := m["a"]
	fmt.Println(name, ok)
	fmt.Println(m)

	delete(m, "a")
	name, ok = m["a"]
	fmt.Println(name, ok)
	fmt.Println(m)
}
