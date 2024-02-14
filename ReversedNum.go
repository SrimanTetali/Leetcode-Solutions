package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var reversedNumber int
	for x != 0 {
		Remainder := x % 10
		x /= 10
		reversedNumber = reversedNumber*10 + Remainder
	}
	if reversedNumber > math.MaxInt32 || reversedNumber < math.MinInt32 {
		return 0
	}
	return reversedNumber
}

func main() {
	var input int
	fmt.Print("Enter the integer to reverse: ")
	fmt.Scan(&input)
	fmt.Println(reverse(input))
}
