package main

import "fmt"

func main() {
	mp := map[string]string{
		"a": "aaaaaa",
		"b": "bbbbbb",
		"c": "cccccc",
	}

	ss := []string{"a", "c", "d"}

	if sliceElemIsMapKey(mp, ss) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}

func sliceElemIsMapKey(mp map[string]string, slice []string) bool {
	for _, elem := range slice {
		if _, ok := mp[elem]; !ok {
			return false
		}
	}
	return true
}
