package main

import "fmt"

func plusMinus(arr []int32) {
	var positive float64
	var negative float64
	var zero float64
	arrLenght := len(arr)
	for _, i := range arr {
		if i > 0 {
			positive++
		} else if i < 0 {
			negative++
		} else if i == 0 {
			zero++
		}
	}
	fmt.Printf("%0.6f\n", float64(positive)/float64(arrLenght))
	fmt.Printf("%0.6f\n", float64(negative)/float64(arrLenght))
	fmt.Printf("%0.6f\n", float64(zero)/float64(arrLenght))
}

func main() {
	plusMinus([]int32{10})
}
