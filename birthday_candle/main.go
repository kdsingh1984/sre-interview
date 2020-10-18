package main

import "fmt"

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	countMap := make(map[int32]int32)
	var maxNum int32
	for _, i := range candles {
		countMap[i]++
		if i > maxNum {
			maxNum = i
		}
	}
	return countMap[maxNum]
}

func main() {
	fmt.Println(birthdayCakeCandles([]int32{3, 3, 1, 4, 5}))
}
