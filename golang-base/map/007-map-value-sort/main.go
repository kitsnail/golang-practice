//要对golang map按照value进行排序，思路是直接不用map，用struct存放key和value，实现sort接口，就可以调用sort.Sort进行排序了。
// A data structure to hold a key/value pair.

package main

import (
	"fmt"
	"sort"
)

func main() {
	// To create a map as input
	m := make(map[string]int)
	m["b"] = 2
	m["a"] = 0
	m["c"] = 1
	pl := sortMapByValue(m)
	for _, p := range pl {
		fmt.Println("Key:", p.Key, "  ", "value:", p.Value)
	}
}

type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}
