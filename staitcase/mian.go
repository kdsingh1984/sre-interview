package main

import "fmt"

// func stepPerms(n int32) int32 {
// 	if n == int32(1) {
// 		return n
// 	} else if n == 2 || n == 3 {
// 		return int32(2) * stepPerms(int32(n-1))
// 	}
// 	return stepPerms(n-1) + stepPerms(n-2) + stepPerms(n-3)
// }

// func stepPerms(n int32) int32 {
//     if n == int32(1) {
//         return int32(1)
//     } else if n == 2 {
//         return int32(2)
//     } else if n == 3 {
//         return int32(4)
//     }
//     return stepPerms(n-1) + stepPerms(n-2) + stepPerms(n-3)
// }

var countMap map[int32]int32 = make(map[int32]int32)

func stepPerms(n int32) int32 {
	var stepn1 int32
	var stepn2 int32
	var stepn3 int32
	if n == int32(1) {
		return int32(1)
	} else if n == 2 {
		return int32(2)
	} else if n == 3 {
		return int32(4)
	}
	if n1Value, ok := countMap[n-1]; !ok {
		countMap[n-1] = stepPerms(n - 1)
		stepn1 = countMap[n-1]
	} else {
		stepn1 = n1Value
	}
	if n2Value, ok := countMap[n-2]; !ok {
		countMap[n-2] = stepPerms(n - 2)
		stepn2 = countMap[n-2]
	} else {
		stepn2 = n2Value
	}
	if n3Value, ok := countMap[n-3]; !ok {
		countMap[n-3] = stepPerms(n - 3)
		stepn3 = countMap[n-3]
	} else {
		stepn3 = n3Value
	}
	return stepn1 + stepn2 + stepn3
}

func main() {
	fmt.Println(stepPerms(35))
}
