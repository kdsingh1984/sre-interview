package main

import "fmt"

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	for i := 1; i < 100000; i++ {
		if x1 > x2 && v1 > v2 {
			return "NO"
		} else if x2 > x1 && v2 > v1 {
			return "NO"
		}
		x1 = x1 + v1
		x2 = x2 + v2
		if x1 == x2 {
			return "YES"
		}
	}
	return "NO"
}

func main() {
	fmt.Println(kangaroo(0, 2, 5, 3))
	fmt.Println(kangaroo(0, 3, 4, 2))
	fmt.Println(kangaroo(2081, 8403, 9107, 8400))
}
