package main

import (
	"fmt"
	"math"
)

func sumofArray(arr []int32, ignoreIndex int) int64 {
	var sum int64
	for i := 0; i < len(arr); i++ {
		if i != ignoreIndex {
			sum += int64(arr[i])
		}
	}
	return sum
}

func miniMaxSum(arr []int32) {
	var min float64 = math.Pow(50, 9)
	var max int64
	for i := 0; i < len(arr); i++ {
		sum := sumofArray(arr, i)
		if sum > max {
			max = sum
		}
		if float64(sum) < min {
			min = float64(sum)
		}
	}
	fmt.Printf("%.0f %d\n", min, max)
}

func main() {
	//fmt.Println(sumofArray([]int32{256741038, 623958417, 467905213, 714532089, 938071625}, 0))
	miniMaxSum([]int32{256741038, 623958417, 467905213, 714532089, 938071625})
}
