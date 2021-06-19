package main

import "fmt"

func duplicateInArray(arr []int) int {
	s := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				s = append(s, arr[i])
				return s[0]
			}
		}
	}
	return -1
}

//Otra Solucion
/* func duplicateInArray(arr []int) int{
	visited := make(map[int]bool)
	for i:=0; i<len(arr); i++{
		if visited[arr[i]] == true{
			return arr[i]
		} else {
			visited[arr[i]] = true
		}
	}
	return -1
} */

func main() {
	fmt.Println(duplicateInArray([]int{1, 2, 7, 2, 4}))
	fmt.Println(duplicateInArray([]int{1, 4, 7, 2, 3}))
}
