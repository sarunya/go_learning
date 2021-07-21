package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	xs := fmt.Sprint(x)
	xsl := len(xs) - 1

	for i := 0; i < (xsl+1)/2; i += 1 {
		fmt.Println(i, xsl-i)
		if xs[i] != xs[xsl-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(12321))
}
