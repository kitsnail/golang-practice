package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	timeLayout   = "2006-01-02 15:04:05"
	durationTime = float64(24 * 3600)
)

func main() {
	cname := "para.csv"
	MergeDataByDay(cname, durationTime)
}

// merge data by same day
func MergeDataByDay(path string, duration float64) {
	var mlist = []string{}
	var count = 0

	f, err := os.Open(path)
	if err != nil {
		log.Fatalln("open csv file error:", err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Read()
	var list []string
	for {
		list, err = r.Read()
		if err != nil {
			if err != io.EOF {
				log.Fatalln("read csv file error:", err.Error())
			} else {
				// 当读到最后io.EOF则计算平均值
				// 计算平均值
				alist := avgList(mlist, count)
				fmt.Println(alist[2])
				fmt.Println(len(alist))
			}
			break
		}
		if len(mlist) == 0 {
			mlist = append(mlist, list...)
			count++
			continue
		}

		// 判断是否是最后一行记录

		// 日期比较，如果mlist，list的是同一天的数据则进行合并操作
		// 否则把mlist的数据输出，然后重置mlist，并把新的list数据
		// 赋值给mlist
		if !isSameDate(mlist[0], list[0]) {
			// 输出mlist
			fmt.Println("计算平均值...")
			// 计算平均值
			alist := avgList(mlist, count)
			fmt.Println(alist)
			// 重置mlist，清空mlist
			mlist = []string{}
			//list赋值给mlist
			mlist = append(mlist, list...)
			continue
		}

		// 合并mlist与list的每个一一对应的指标项
		// 这里求出指标项的范围包括集群里的所有节点
		for i := 1; i < len(mlist)-1; i++ {
			mitem, err := strconv.ParseFloat(mlist[i], 5)
			if err != nil {
				break
			}
			litem, err := strconv.ParseFloat(list[i], 5)
			if err != nil {
				break
			}

			mlist[i] = fmt.Sprintf("%.5f", mitem+litem)
		}
		count++
	}

}

// 计算列表数据的平均值
func avgList(list []string, count int) (ret []string) {
	for i, l := range list {
		if i == 0 {
			ret = append(ret, l)
			continue
		}
		item, err := strconv.ParseFloat(l, 5)
		if err != nil {
			ret = append(ret, "")
		}
		ret = append(ret, fmt.Sprintf("%.5f", item/float64(count)))
	}
	return ret
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

// 默认s1,s2的时间格式: 2006-01-02 15:04:05
func isSameDate(s1, s2 string) bool {
	return strings.Fields(s1)[0] == strings.Fields(s2)[0]
}

// s1,s2 是时间字符串
// layout 时间格式
// dur 默认是时间单位s
func CompareTimesOverRange(s1, s2 string, layout string, dur float64) (bool, error) {
	t1, err := time.Parse(layout, s1)
	if err != nil {
		return false, err
	}
	t2, err := time.Parse(layout, s2)
	if err != nil {
		return false, err
	}
	return compareTimesOverRange(t1, t2, dur), nil
}

func compareTimesOverRange(t1, t2 time.Time, duration float64) bool {
	dt := t2.Sub(t1).Seconds()
	if dt > duration {
		return false
	}
	return true
}
