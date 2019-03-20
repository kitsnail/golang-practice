package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	url := "https://file.paratera.com/api/file/manager/ls?path=pcs://path.tianhe2-C.GUANGZHOU/paratera_60/home/wanghui/matlab.zip"
	//url := "https://file.paratera.com/api/file/stream/download?path=pcs://path.tianhe2-C.GUANGZHOU/paratera_60/home/wanghui/test.small"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("PARA_TOKEN", `4e1EwrjrtRHEdiZrZPUDgN3UJ9sL1gcKcYhenUkofpM-2554739209`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
