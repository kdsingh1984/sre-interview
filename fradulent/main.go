package main

import (
	"fmt"
)

type int32arr []int32

func (in int32arr) Len() int {
	return len(in)
}

func (in int32arr) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func (in int32arr) Less(i, j int) bool {
	return in[i] < in[j]
}

func median(arr []int32, arrLen int32) float64 {
	// fmt.Printf("Length of arr is %v", arrLen)
	var m float64
	if arrLen%2 == 0 {
		n := arrLen / 2
		m = float64(arr[n]+arr[n-1]) / float64(2.0)
		// fmt.Printf("meidan is %v, upper: %v, lower: %v", m, float64(arr[n]+arr[n-1]), float64(2.0))
	} else if arrLen%2 != 0 {
		n := (arrLen - 1) / 2
		m = float64(arr[n])
	}
	return m
}

func countSort(arr []int32) []int32 {
	var maxvalue int32
	for _, i := range arr {
		if i > maxvalue {
			maxvalue = i
		}
	}
	var countArr = make([]int, maxvalue+1)
	for i := 0; i < len(arr); i++ {
		countArr[arr[i]]++
	}
	for i := 1; i < len(countArr); i++ {
		countArr[i] = countArr[i] + countArr[i-1]
	}
	var newArr = make([]int32, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		countArr[arr[i]]--
		newArr[countArr[arr[i]]] = arr[i]
	}
	return newArr
}

func activityNotifications(expenditure []int32, d int32) int32 {
	totalalerts := int32(0)
	var sortedArr []int32
	var newArr []int32
	if int32(len(expenditure)) > d {
		newArr := expenditure[0:d]
		sortedArr = countSort(newArr)
	}
	for i := d; i < int32(len(expenditure)); i++ {
		//newArr := expenditure[i-d : i]
		// sortedArr := int32arr(newArr)
		// sort.Sort(sortedArr)
		median := median(sortedArr, d)
		newArr = sortedArr[1:]
		newArr = append(newArr, expenditure[i])
		sortedArr = countSort(newArr)
		// fmt.Println(sortedArr)
		// fmt.Println("median is ", median)
		if float64(expenditure[i]) >= 2*median {
			totalalerts++
		}
	}
	return totalalerts
}

func main() {
	//fmt.Println(countSort([]int32{2, 3, 4, 2, 3}))
	fmt.Println(activityNotifications([]int32{2, 3, 4, 2, 3, 6, 8, 4, 5, 2, 3, 4, 2, 3, 6, 8, 4, 5, 2, 3, 4, 2, 3, 6, 8, 4, 5, 2, 3, 4, 2, 3, 6, 8, 4, 5, 2, 3, 4, 2, 3, 6, 8, 4, 5, 2, 3, 4, 2, 3, 6, 8, 4, 5, 2, 3, 4, 2, 3, 6, 8, 4, 5, 2, 3, 4, 2, 3, 6, 8, 4, 5, 2, 3, 4, 2, 3, 6, 8, 4, 5}, 5))
	fmt.Println(activityNotifications([]int32{1, 2, 3, 4, 4}, 4))
}
