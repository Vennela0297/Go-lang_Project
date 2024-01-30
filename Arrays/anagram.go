package main

import(
	"fmt"
	"sort"
) 

func isAnagram(str1, str2 string) bool {
	
	runes1 := []rune(str1)
	runes2 := []rune(str2)

	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

	return string(runes1) == string(runes2)
}

func main() {
	str1 := "listen"
	str2 := "silent"

	if isAnagram(str1, str2) {
		fmt.Println(str1, "and", str2, "are anagrams")
	} else {
		fmt.Println(str1, "and", str2, "are not anagrams")
	}
}