package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	raw := "pcs://path.tianhe2-C.GUANGZHOU/paratera_60/home/wanghui/pcloud/papp2/cloud2/pc1"
	u, err := url.Parse(raw)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.RequestURI())
	fmt.Println(u.EscapedPath())
}
