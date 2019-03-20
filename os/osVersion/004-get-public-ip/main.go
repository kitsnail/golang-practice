package main

import (
	"fmt"
	"io/ioutil"
   "os"
   "net/http"
)

func main(){
   ip,err := get_external()
   if err != nil {
	   fmt.Println(err)
	   os.Exit(1)
   }
   fmt.Print(ip)
}

func get_external() (ip string ,err error) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	return string(body),nil
}
