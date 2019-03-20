package main

import (
	"fmt"
)

func main() {
	mmp := map[string]string{
		"node1": "nfs1",
		"node2": "nfs2",
		"node3": "nfs2",
		"node4": "nfs1",
		"node5": "nfs1",
	}
	mps := mapkey2value(mmp)
	fmt.Println(mps)
}

func mapkey2value(mmp map[string]string) (ret map[string][]string) {
	ret = make(map[string][]string)
	for k, v := range mmp {
		ret[v] = append(ret[v], k)
	}
	return
}
