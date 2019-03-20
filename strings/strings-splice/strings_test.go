package main

import "testing"

func StringPlus() string {
	var s string
	s += "昵称" + ":" + "飞雪无情" + "\n"
	s += "博客" + ":" + "http://www.flysnow.org/" + "\n"
	s += "微信公众号" + ":" + "flysnow_org"
	return s
}

func BenchmarkStringPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringPlus()
	}
}
