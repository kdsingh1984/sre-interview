package main

import "fmt"

func sum(worklist []int32, startindex int32, firstskip int32) int32 {
	var initialsum int32
	initialsum = worklist[startindex] + worklist[startindex+firstskip+1]
	for i := startindex + firstskip + 1; i <= int32(len(worklist)); i += 2 {
		initialsum += worklist[i]
	}
	return initialsum
}

//[1,2,3,4,5,6,7,8,9,10]
//[3 5 -7 8 10]
func maxSubsetSum(arr []int32) int32 {
	var sum int32
	for i := 0; i < len(arr); i++ {

	}
}

func main() {
	fmt.Println(maxSubsetSum([]int32{3, 5, -7, 8, 10}))
}
