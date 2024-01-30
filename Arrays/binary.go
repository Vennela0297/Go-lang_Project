package main

import (
	"fmt"
)

func binarySearch(slice []int, target int) int {
	X := 0
	Y := len(slice) - 1

	for X <= Y {
		a := (X + Y) / 2
		if slice[a] == target {
			return a 
		} else if slice[a] < target {
			X = a + 1 
		} else {
			Y = a - 1 
		}
	}
	return -1 
}

func main() {
	
	slice := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 11

	index := binarySearch(slice, target)

	if index != -1 {
		fmt.Println(target, index)
	} else {
		fmt.Println(target)
	}
}