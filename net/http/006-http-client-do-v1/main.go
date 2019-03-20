package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	/*
		f, err := os.OpenFile("cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		log.Println("CPU Profile started")
		pprof.StartCPUProfile(f)
	*/

	defaultEndpoint := "https://file.paratera.com"
	filename := "matlab.zip"
	targeturl := fmt.Sprintf("path=pcs://path.tianhe2-C.GUANGZHOU/paratera_60/home/wanghui/%s", filename)

	url := fmt.Sprintf("%s/api/file/manager/ls?%s", defaultEndpoint, targeturl)

	token := "4e1EwrjrtRHEdiZrZPUDgN3UJ9sL1gcKcYhenUkofpM-2554739209"
	size, err := getFileSize(url, token)
	//size, err := getFileSize(url, token)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(size)

	//pprof.StopCPUProfile()

}

func cpuProfile() {

	time.Sleep(60 * time.Second)
	fmt.Println("CPU Profile stopped")
}

func newHTTPClient() *http.Client {
	return &http.Client{}
}

/*
func getFileSize(url string) (int64, error) {
	tb := time.Now()
	fmt.Println(tb)
	client := newHTTPClient()
	fmt.Printf("new client: %.2f\n", time.Since(tb).Seconds())
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	fmt.Printf("new request: %.2f\n", time.Since(tb).Seconds())
	//req.Header.Set("PARA_TOKEN", token)
	//fmt.Printf("set Header: %.2f\n", time.Since(tb).Seconds())

	resp, err := client.Do(req)
	fmt.Printf("client do: %.2f\n", time.Since(tb).Seconds())
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	size := resp.ContentLength
	fmt.Printf("resp ContentLength: %.2f\n", time.Since(tb).Seconds())
	return size, nil
}
*/

func getFileSize(url string, token string) (string, error) {
	tb := time.Now()
	fmt.Println(tb)
	client := &http.Client{}
	fmt.Printf("new client: %.2f\n", time.Since(tb).Seconds())
	fmt.Println(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	//req, err := http.NewRequest(http.MethodGet, "https://file.paratera.com/api/file/manager/ls?path=pcs://path.tianhe2-C.GUANGZHOU/paratera_60/home/wanghui/matlab.zip", nil)
	if err != nil {
		return "", err
	}

	fmt.Printf("new request: %.2f\n", time.Since(tb).Seconds())
	req.Header.Add("PARA_TOKEN", token)
	fmt.Printf("set Header: %.2f\n", time.Since(tb).Seconds())

	resp, err := client.Do(req)
	fmt.Printf("client do: %.2f\n", time.Since(tb).Seconds())
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("resp ContentLength: %.2f\n", time.Since(tb).Seconds())
	return string(body), nil
}
