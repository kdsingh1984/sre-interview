package main

import (
	"fmt"
	"strings"
)

func staircase(n int32) {
	var returnList []string
	for i := int32(1); i <= n; i++ {
		inputStr := fmt.Sprintf("%s%s", strings.Repeat(" ", int(n-i)), strings.Repeat("#", int(i)))
		returnList = append(returnList, inputStr)
	}
	fmt.Println(strings.Join(returnList, "\n"))
}

func main() {
	staircase(10)
}
