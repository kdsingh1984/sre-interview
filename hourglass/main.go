package main

// 1 1 1 0 0 0
// 0 1 0 0 0 0
// 1 1 1 0 0 0
// 0 0 2 4 4 0
// 0 0 0 2 0 0
// 0 0 1 2 4 0

//[[1,1,1,0,0,0],[0,1,0,0,0,0],[1,1,1,0,0,0],[0,0,2,4,4,0],[0,0,0,2,0,0],[0,0,1,2,4,0]]


func hourglassSum(arr [][]int32) int32 {
	var maxSum int32
	for i := 0; i <= 3; i++ {
		arrSet := arr[i : i+3]
		for j := 1; j <= 4; j++ {
			s := arrSet[0][j-1] + arrSet[0][j] + arrSet[0][j+1] + arrSet[1][j] + arrSet[2][j-1] + arrSet[2][j] + arrSet[2][j+1]
			if s > maxSum {
				maxSum = s
			}
		}
	}
	return maxSum
}

func main() {
	fmt.Println(hourglassSum([]int32{[1,1,1,0,0,0],[0,1,0,0,0,0],[1,1,1,0,0,0],[0,0,2,4,4,0],[0,0,0,2,0,0],[0,0,1,2,4,0]})
}
