package main

import (
	"fmt"
	"sort"
	"strings"
)

func sherlockAndAnagrams(s string) int32 {
	stringMap := make(map[string]int)
	stringArr := strings.Split(s, "")
	for i := 0; i < len(stringArr); i++ {
		stringMap[stringArr[i]]++
		for j := i + 2; j < len(stringArr)+1; j++ {
			b := strings.Split(s, "")
			a := b[i:j]
			sort.Strings(a)
			sortedString := strings.Join(a, "")
			stringMap[sortedString]++
		}
	}
	sum := 0
	for _, v := range stringMap {
		sum += (v * (v - 1)) / 2
	}
	return int32(sum)
}

func main() {
	fmt.Println(sherlockAndAnagrams("kkkk"))
}
