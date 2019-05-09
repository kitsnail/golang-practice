package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const bufferSize = 100

func main() {
	url := "http://66.42.91.35/textn"
	tsize, err := httpGetFileSize(url)
	if err != nil {
		log.Fatalln(err)
	}
	m := New(httpGetBody)
	var wg sync.WaitGroup

	for _, rang := range createRanges(tsize, bufferSize) {
		wg.Add(1)
		go func(url string, rang Rang) {
			_, err := m.Get(url, rang)
			if err != nil {
				log.Println(err)
			}
			wg.Done()
		}(url, rang)
	}
	wg.Wait()

	for k, v := range m.cache {
		fmt.Println(k)
		fmt.Println(v.res.value)
	}
}

type Rang struct {
	Begin int64
	End   int64
}

func createRanges(totalSize int64, bs int64) (ranges []Rang) {
	var begin int64
	var end int64

	for begin < totalSize {
		end += bs
		ranges = append(ranges, Rang{begin, end})
		begin = end
	}
	ranges[len(ranges)-1].End = totalSize - 1
	return
}

func httpGetFileSize(url string) (int64, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	header := resp.Header
	clength := header.Get("Content-Length")
	return strconv.ParseInt(clength, 10, 64)
}

func httpGetBody(url string, rang Rang) (interface{}, error) {
	c := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rang.Begin, rang.End))
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

type Func func(key string, arg Rang) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (m *Memo) Get(key string, arg Rang) (value interface{}, err error) {
	keys := fmt.Sprintf("%d-%d", arg.Begin, arg.End)
	m.mu.Lock()
	e := m.cache[keys]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		m.cache[keys] = e
		m.mu.Unlock()

		e.res.value, e.res.err = m.f(key, arg)
		close(e.ready)
	} else {
		m.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}
