package main

import "fmt"

[]int32{100, 90, 90, 80, 75, 60}, []int32{50, 65, 77, 90, 102}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	returnList := make([]int32, len(player))
	rankedMap := make(map[int32]int32)
	rankNum := int32(1)
	for _, rank := range ranked {
		if _, ok := rankedMap[rank]; !ok {
			rankedMap[rank] = rankNum
			rankNum++
		}
	}
	for i := 0; i < len(player); i++ {
		if player[i] > ranked[0] {
			returnList[i] = 1
			continue
		} else if player[i] < ranked[len(ranked)-1] {
			returnList[i] = int32(len(rankedMap) + 1)
			continue
		}
		for j := 0; j < len(ranked); j++ {
			if player[i] >= ranked[j] {
				returnList[i] = rankedMap[ranked[j]]
				break
			}
		}
	}
	return returnList
}

func main() {
	fmt.Println(climbingLeaderboard([]int32{100, 90, 90, 80, 75, 60}, []int32{50, 65, 77, 90, 102}))
}
