package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	GitlabEdpoint      = "https://git.paratera.net/api/v4"
	GitlabPrivateToken = "7qcmkuZxBmHP-JxMy5xx"
)

func main() {
	url := fmt.Sprintf("%s/groups/sre/issues?", GitlabEdpoint)
	GetByToken(url, GitlabPrivateToken)
}

// read response 'header' and 'body' content
func GetByToken(url string, token string) {
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
	}
	client := http.Client{Transport: tr}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("PRIVATE-TOKEN", token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	header := resp.Header
	for k, v := range header {
		fmt.Println(k, ":", v)
	}

	// parse header 'Link' value
	links := header["Link"]
	for i, v := range links {
		fmt.Printf("%d:%s\n", i, v)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
