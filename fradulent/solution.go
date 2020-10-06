package main

import "sort"

type foo []int32

func (f foo) Len() int {
	return len(f)
}

func (f foo) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f foo) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func sortList(int32List []int32) []int32 {
	a := foo(int32List)
	sort.Sort(a)
	return a
	// var intList []int
	// for _, i := range int32List {
	//     intList = append(intList, int(i))
	// }
	// sort.Ints(intList)
	// return intList
}

func median(spendList []int32) int32 {
	var m int32
	listLen := len(spendList)
	if listLen%2 == 0 {
		element := listLen / 2
		m = (spendList[element-1] + spendList[element]) / 2
	} else if listLen%2 != 0 {
		element := (listLen - 1) / 2
		m = spendList[element]
	}
	return m
}

// Complete the activityNotifications function below.
func activityNotifications(expenditure []int32, d int32) int32 {
	totalAlerts := int32(0)
	finalExpenseElement := int32(len(expenditure)) - d - 1
	for i := int32(0); i < finalExpenseElement; i++ {
		sortedList := sortList(expenditure[i : i+d])
		m := median(sortedList)
		nextElement := expenditure[i+d]
		if 2*int32(m) >= nextElement {
			totalAlerts += 1
		}
	}
	return totalAlerts
}
