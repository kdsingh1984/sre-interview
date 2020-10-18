package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	var multipleofFiveArray []int32
	for i := int32(1); i <= 20; i++ {
		multipleofFiveArray = append(multipleofFiveArray, i*5)
	}
	for gradeIndex := 0; gradeIndex < len(grades); gradeIndex++ {
		var nextFiveMultiple int32
		if grades[gradeIndex] >= 38 && grades[gradeIndex] < 100 {
			for _, i := range multipleofFiveArray {
				if i > grades[gradeIndex] {
					nextFiveMultiple = i
					break
				}
			}
			if nextFiveMultiple-grades[gradeIndex] < 3 {
				grades[gradeIndex] = nextFiveMultiple
			}
		}
	}
	return grades
}

func main() {
	fmt.Println(gradingStudents([]int32{73, 67, 38, 33, 23, 39, 100}))
}
