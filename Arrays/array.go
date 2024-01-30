package main

import
(
	"fmt"
)
	
func main(){

	var a [] int = []int{1,22,33,4,22,5,3,8,33,22}
	var b int
	b = 0

	for i := 0; i < len(a); i++ {

		for j := 0; j < b; j++{
			if a[i] == a[j]{

				fmt.Println(a[i])

			}
		}
		b = b+1
		
	}

}

