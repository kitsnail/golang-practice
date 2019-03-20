package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	client "github.com/influxdb/influxdb/client/v2"
)

const (
	Tlayout  = "2006-01-02 15:04:05"
	MyDB     = "mydb"
	username = ""
	password = ""
)

func main() {
	var hostsMetric = make(map[string]map[string]int)
	var metrics map[string]int
	var clusterName string

	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// csv 文件读取解析
	input := "para-input.csv"
	f, err := os.Open(input)
	checkError("Cannot open file", err)
	defer f.Close()
	r := csv.NewReader(f)
	keys, err := r.Read()
	checkError("Cannot read file", err)

	// 生成集群主机节点指标项的字典
	for i, key := range keys {
		if i == 0 {
			continue
		}
		if len(key) == 0 {
			continue
		}
		slice1 := strings.Split(key, "/")
		clusterName = slice1[0]
		hostname := slice1[1]
		item := slice1[2]

		if _, ok := hostsMetric[hostname]; !ok {
			metrics = make(map[string]int)
			hostsMetric[hostname] = make(map[string]int)
		}

		metrics[item] = i
		hostsMetric[hostname] = metrics
	}

	for {
		list, err := r.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println("读取文件内容失败,错误信息:", err.Error())
			}
			break
		}

		// 按主机节点一次生成influxdb的mesurement

		for host, metrics := range hostsMetric {
			// 创建tags
			tags := map[string]string{
				"host":    host,
				"cluster": clusterName,
			}
			fields := make(map[string]interface{})

			for item, number := range metrics {
				fields[item] = list[number]
			}

			ts, err := time.Parse(Tlayout, list[0])
			if err != nil {
				log.Println(err)
				continue
			}

			pt, err := client.NewPoint("paramonMetric3", tags, fields, ts)
			if err != nil {
				log.Println(err)
				continue
			}

			bp.AddPoint(pt)

			fmt.Println(bp)

		}
	}
	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}

}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
