package main

import "fmt"

func main() {
	s1 := []string{"1", "2", "3", "4", "7"}
	s2 := []string{"4", "5", "7", "8"}
	mmp := map[string][]string{
		"foo": s1,
		"bar": s2,
	}
	fmt.Println(mmp)
}

func mergeElemMap(mp map[string][]string) (ret map[string][]string) {
	ret = make(map[string][]string)
	for k, v := range mp {
		elm, ok := ret[k]
		if !ok {
			ret[k] = v
		}

	}
}
