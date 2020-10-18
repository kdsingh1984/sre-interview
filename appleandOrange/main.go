package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var applesCount int32
	var orangesCount int32
	for _, apple := range apples {
		if apple+a >= s && apple+a <= t {
			applesCount++
		}
	}
	for _, orange := range oranges {
		if orange+b >= s && orange+b <= t {
			orangesCount++
		}
	}
	fmt.Println(applesCount)
	fmt.Println(orangesCount)
}

func main() {
	countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
}
