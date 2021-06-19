package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	n := x
	sol := 0
	if x < 0 {
		return false
	} else {
		for n > 0 {
			ultimo := n % 10
			sol = (sol * 10) + ultimo
			n /= 10
		}
		if sol == x && sol < math.MaxInt32 {
			return true
		} else {
			return false
		}
	}
}

func main() {
	x := 10
	fmt.Println(isPalindrome(x))
}
