package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	cname := "para.csv"
	qKeys := []string{"cpu_all_cpi",
		"cpu_all_flops",
		"cpu_all_hi",
		"cpu_all_idle",
		"cpu_all_ips",
		"cpu_all_llcmiss",
		"cpu_all_ni",
		"cpu_all_si",
		"cpu_all_sy",
		"cpu_all_sywa",
		"cpu_all_us",
		"cpu_all_wa",
		"cpu_core_top_flops",
		"cpu_top_flops",
		"disk_all_read",
		"disk_all_write",
		"ftime",
		"infiniband_all_rcv",
		"infiniband_all_xmt",
		"mem_all_memRatio",
		"mem_all_swapRatio",
		"mem_swap_in",
		"mem_swap_out",
		"net_all_recv",
		"net_all_send",
		"net_num",
		"nfs_all_in",
		"nfs_all_out",
		"nfs_pkg_in",
		"nfs_pkg_out"}

	mp := readColLists(cname, qKeys)
	for k, v := range mp {
		fmt.Println(k, ":", v)
		fmt.Println("\n")
	}
}

func readColLists(path string, metric []string) (metrMap map[string][]int) {
	metrMap = make(map[string][]int)

	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	r := csv.NewReader(f)
	heads, err := r.Read()
	if err != nil {
		log.Fatalln(err)
	}

	for _, key := range metric {
		cols := []int{}
		metrMap[key] = cols
		for i, h := range heads {
			s1 := strings.Split(h, "/")
			if len(s1) == 3 {
				if strings.Compare(s1[2], key) == 0 {
					cols = append(cols, i)
				}
			}
		}
		metrMap[key] = cols
	}
	return metrMap
}
