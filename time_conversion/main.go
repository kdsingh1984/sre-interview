package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	timeArry := strings.Split(s, "")
	ampm := strings.Join(timeArry[8:10], "")
	hour := strings.Join(timeArry[0:2], "")
	if ampm == "AM" {
		if hour == "12" {
			return "00" + s[2:8]
		}
		return strings.Join(timeArry[0:8], "")
	}
	var pmHour string
	if hour == "12" {
		pmHour = hour
	} else if hour != "12" {
		hourInt, _ := strconv.Atoi(hour)
		pmHour = strconv.Itoa(hourInt + 12)
	}
	return string(pmHour) + s[2:8]
}

func main() {
	fmt.Println(timeConversion("12:05:45AM"))
}
