package main

import
(
	"fmt"
)
	
func main(){

	var a [] int = []int{1,2,33,44,54,34,54,10,33,22}
	var b int
	b = 0

	for i := 0; i < len(a); i++ {

		for j := 0; j < len(a); j++{
			if j == i {
				
				continue
			} 
			if a[i] + a[j] <= 10{

				b = b+1
				fmt.Println(i,j)
				//fmt.Println(a[i], a[j])

			} 
			
			}
			
			}
			if b == 0{
				fmt.Println(-1)
			}
		}
		
		
	



