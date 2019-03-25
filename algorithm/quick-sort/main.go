package main

import (
    "fmt"
)

func main(){
    list := []int{9,4,2,49,20,48,98,3,1,8,10}
	fmt.Println("old:",list)
	qsort(list)
	fmt.Println("new:",list)

}

func qsort(data []int){
    if len(data) <= 1 {
	   return
	}

	mid := data[0]
	head,tail := 0,len(data)-1
	for i :=1; i <= tail; {
	    if data[i] > mid {
		    data[i],data[tail] = data[tail],data[i]
			tail --
		}else {
		    data[i],data[head] = data[head],data[i]
			head ++
			i ++
		}
	}
	qsort(data[:head])
	qsort(data[head+1:])
}
