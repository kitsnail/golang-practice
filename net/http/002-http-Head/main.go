package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

const (
	GitlabEdpoint      = "https://git.paratera.net/api/v4"
	GitlabPrivateToken = "7qcmkuZxBmHP-JxMy5xx"
)

func main() {
	url := fmt.Sprintf("%s/groups/sre/issues?", GitlabEdpoint)
	HeadByToken(url, GitlabPrivateToken)
}

func HeadByToken(url string, token string) {
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
	}
	client := http.Client{Transport: tr}

	req, err := http.NewRequest(http.MethodHead, url, nil)
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
	fmt.Println(header)
}
