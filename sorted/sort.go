package main

import (
	"fmt"
	"sort"
)

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

func main() {
	a := []int32{3, 2, 1}
	b := foo(a)
	f := foo{3, 2, 1}
	sort.Sort(f)
	fmt.Println(f)
	sort.Sort(b)
	fmt.Println(b)
}
