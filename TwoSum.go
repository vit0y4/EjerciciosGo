package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var sol []int
	for n := 0; n < len(nums); n++ {
		for i := n + 1; i < len(nums); i++ {
			if (nums[n] + nums[i]) == target {
				sol = []int{n, i}
				return sol
			}
		}
	}
	return sol
}

func main() {
	array := []int{2, 5, 5, 11}
	fmt.Println(array)
	fmt.Println(twoSum(array, 10))

}
