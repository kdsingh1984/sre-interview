package main

import "fmt"

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	matrixSize := len(arr)
	var leftToRight int32
	var rightToLeft int32
	for i := 0; i < matrixSize; i++ {
		leftToRight += arr[i][i]
		rightToLeft += arr[i][matrixSize-1-i]
	}
	if leftToRight > rightToLeft {
		return leftToRight - rightToLeft
	}
	return rightToLeft - leftToRight
}

func main() {
	arr := [][]int32{[]int32{11, 2, 4}, []int32{4, 5, 6}, []int32{10, 8, -12}}
	fmt.Println(diagonalDifference(arr))
}
