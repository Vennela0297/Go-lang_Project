package main

import(
	"fmt"
)

func removeElement(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:]) 
	return slice[:len(slice)-1]          
}

func main() {
	
	slice := []int{1, 2, 3, 4, 5}
	indexToRemove := 2
	slice = removeElement(slice, indexToRemove)

	fmt.Println(slice)
}