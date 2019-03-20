package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//cname := "para.csv"
	cname := "output.csv"
	lines := []int{2}
	FromLineRead(lines, cname)
}

func FromLineRead(lines []int, path string) {
	File, err := os.Open(path)
	if err != nil {
		log.Println("读取csv文件失败:", err.Error())
		return
	}
	defer File.Close()
	r := csv.NewReader(File)
	r.Read()
	var list []string
	var line, index int
	for {
		list, err = r.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println("读取文件内容失败,错误信息:", err.Error())
			}
			break
		}
		if lines[index] == line {
			//fmt.Println(list)
			fmt.Println(list[0])
			fmt.Println(len(list))

			index++
			if index >= len(lines) {
				break
			}
		}
		line++
	}
}

//func ListSum()
