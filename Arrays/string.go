package main

import(
	"fmt"
)
func main(){
	var a string = "vennela"

	//a.length();

	myMap := make(map[string]int)

	for i := 0; i < len(a); i++{
		char := string(a[i])

		value, exists := myMap[char]

		if exists {
			
			myMap[char] = value + 1
		} else {	
			
			myMap[char] = 1
			//fmt.Println(myMap)
		}
		
	}
	var maxChar []string 
	maxCount := 0

	for char, count := range myMap {
		if count > maxCount {

			maxChar = []string{char}
			maxCount = count
		} else if count == maxCount{
			maxChar = append(maxChar, char)
		}
	}
	fmt.Println(maxChar)

}