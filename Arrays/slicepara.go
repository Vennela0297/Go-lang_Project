package main

import (
	"fmt"
)

func rotateLeft(slice []int, n int) []int {
	length := len(slice)
	if length == 0 {
		return slice
	}

	n = n % length

	result := make([]int, length)
	for i := range slice {

		newIndex := (i - n + length) % length
		result[newIndex] = slice[i]
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	rotateBy := 3

	fmt.Println("Original Slice:", slice)

	rotatedSlice := rotateLeft(slice, rotateBy)

	fmt.Println("Rotated Slice:", rotatedSlice)
}
