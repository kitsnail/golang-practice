package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	m := New(httpGetBody)
	var wg sync.WaitGroup
	for _, url := range incomingURLs() {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func incomingURLs() []string {
	return []string{"http://66.42.91.35/text0",
		"http://66.42.91.35/text1",
		"http://66.42.91.35/text2",
		"http://66.42.91.35/text3",
		"http://66.42.91.35/text4",
		"http://66.42.91.35/text5",
		"http://66.42.91.35/text6",
		"http://66.42.91.35/text7",
		"http://66.42.91.35/text8",
		"http://66.42.91.35/text9",
		"http://66.42.91.35/text10"}

}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (m *Memo) Get(key string) (interface{}, error) {
	m.mu.Lock()
	res, ok := m.cache[key]
	if !ok {
		res.value, res.err = m.f(key)
		m.cache[key] = res
	}
	m.mu.Unlock()
	return res.value, res.err
}
