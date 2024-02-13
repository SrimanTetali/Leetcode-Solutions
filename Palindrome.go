package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var revnum int
	for x > revnum {
		revnum = revnum*10 + x%10
		x /= 10
	}
	return x == revnum || x == revnum/10
}
func main() {
	var x int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&x)
	fmt.Print(isPalindrome(x))
}
