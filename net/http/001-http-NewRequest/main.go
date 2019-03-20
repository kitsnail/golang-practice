package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	url = "https://papp.paratera.com/vpn/api/v3/request_areas/GUANGZHOU/putty/0/1"
)

func main() {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	fmt.Println(req)
	if err != nil {
		fmt.Println("1:", err)
		os.Exit(1)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("2:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode, resp.Status)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("3:", err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}
