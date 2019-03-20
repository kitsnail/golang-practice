package main

import (
	"fmt"
)

var (
	user     string
	password string
)

func main() {
	fmt.Println("papp_cloud login test")
	fmt.Printf("username:")
	fmt.Scan(&user)
	//fmt.Printf("password:")
	//fmt.Scan(&password)
	password, err := Getpasswd("input your password:")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n")
	fmt.Printf("Your username is:%s", user)
	fmt.Printf("Your password is:%s", password)
}
