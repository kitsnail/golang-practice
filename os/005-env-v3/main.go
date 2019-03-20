package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	key := "PCLOUD_OLD_PATH"
	value := "/Users/wanghui/Downloads"
	if err := os.Setenv(key, value); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(os.Getenv(key))
	time.Sleep(1 * time.Hour)
}
