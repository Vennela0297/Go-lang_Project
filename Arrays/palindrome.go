package main

func isPalindrome(s string) bool {
	runes := []rune(s)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	input := "adfgc"

	if isPalindrome(input) {
		println(input, "is a palindrome")
	} else {
		println(input, "is not a palindrome")
	}
}