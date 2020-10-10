package main

import "fmt"

// [1,2,2,4], 2

// (0,1,3)
// (0,2,3)

func CombinationCalculator(n int64, r int64) int64 {
	q := int64(1)
	d := int64(1)
	for i := r - 1; i >= int64(0); i-- {
		q *= n - i
		d *= i + 1
	}
	return q / d
}

func countTriplets(arr []int64, r int64) int64 {
	if len(arr) < 3 {
		return 0
	}
	arrMap := make(map[int64]int)
	for _, i := range arr {
		arrMap[i]++
	}
	if r == int64(1) {
		// sum := int64(arrMap[1] - 2)
		// if sum < 0 {
		// 	return 0
		// }
		return CombinationCalculator(int64(arrMap[1]), 3)
	}
	sum := 0
	for k, v := range arrMap {
		second := k * r
		third := second * r
		secondCount, secondOk := arrMap[second]
		thirdCount, thirdOk := arrMap[third]
		if secondOk && thirdOk {
			sum += v * secondCount * thirdCount
		}
	}
	return int64(sum)
}

func main() {
	//fmt.Println(CombinationCalculator(int64(100), int64(3)))
	fmt.Println(countTriplets([]int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1))
	fmt.Println(countTriplets([]int64{2, 4, 8, 16, 2}, 2))
	fmt.Println(countTriplets([]int64{1, 1, 1, 2, 3, 4, 8}, 2))
	// fmt.Println(countTriplets([]int64{1, 1, 1, 1}, 1))
	// fmt.Println(countTriplets([]int64{1, 2, 2, 4}, 2))
	// fmt.Println(countTriplets([]int64{1, 3, 9, 9, 27, 81}, 3))
	// fmt.Println(countTriplets([]int64{1, 5, 5, 25, 125}, 5))
	// fmt.Println(countTriplets([]int64{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 4}, 3))
}
