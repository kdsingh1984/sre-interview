package main

import (
	"fmt"
	"math"
)

func maxSubsetSum(arr []int32) int32 {
	if len(arr) == 1 {
		return arr[0]
	}
	dp := make([]int32, len(arr))
	dp[0] = arr[0]
	dp1float := math.Max(float64(dp[0]), float64(dp[1]))
	dp[1] = int32(dp1float)
	for i := 2; i < len(arr); i++ {
		dpi := math.Max(math.Max(float64(dp[i-1]), float64(arr[i])), float64(dp[i-2]+arr[i]))
		dp[i] = int32(dpi)
	}
	return dp[len(arr)-1]
}

func main() {
	list := []int32{2, 1, 5, 8, 4}
	fmt.Println(maxSubsetSum(list))
}
