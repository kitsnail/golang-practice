package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cb/compress", cbHandler)
	http.HandleFunc("/cb/uncompress", cbHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func cbHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	fmt.Println(string(body))
}
