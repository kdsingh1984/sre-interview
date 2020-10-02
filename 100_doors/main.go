package main

import (
	"fmt"
	"sort"
)

func openDoorsList(doorsStatus map[int]bool) []int {
	returnList := []int{}
	totalDoors := 4
	for skip := 1; skip <= totalDoors; skip++ {
		for door := 1; door <= totalDoors; door += skip {
			doorsStatus[door] = !doorsStatus[door]
		}
	}
	for key, value := range doorsStatus {
		if value {
			returnList = append(returnList, key)
		}
	}
	return returnList
}

func listToMap(inputList []bool) map[int]bool {
	returnMap := make(map[int]bool)
	mapIndex := 1
	for _, i := range inputList {
		returnMap[mapIndex] = i
		mapIndex += 1
	}
	return returnMap
}

func main() {
	doorList := []bool{true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false}
	doorsStatus := listToMap(doorList)
	listlist := openDoorsList(doorsStatus)
	sort.Ints(listlist)
	fmt.Println(listlist)
}
