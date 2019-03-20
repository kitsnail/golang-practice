package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	baseAPI := "http://cloudlab.paratera.com:7777/api/v1"
	ContentType := "application/json;charset=utf-8"
	data := []byte{}
	Login(baseAPI, ContentType, data)
}

type AuthUser struct {
}

//func Login(base string, contentType string, data []byte) (au AuthUser, err error) {
func Login(base string, contentType string, data []byte) {
	url := fmt.Sprintf("%s/userservice/login/", base)
	body := bytes.NewBuffer(data)
	res, err := http.Post(url, contentType, body)
	if err != nil {
		//return nil,err
		log.Fatalln(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatalln(err)
		//	return nil,err
	}
	//return nil
	fmt.Println(result)

}
