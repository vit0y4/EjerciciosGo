package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func longestCommonPrefix(strs []string) string {
	length := len(strs[0])
	word1 := strs[0]
	if len(strs) < 200 {
		for i := 1; i < len(strs); i++ {
			cword := strs[i]
			length = min(length, len(strs[i]))
			for k := 0; k < length; k++ {
				if string(word1[k]) != string(cword[k]) {
					length = k
					break
				}
			}
		}
		if len(strs) == 1 {
			return strs[0]
		} else {
			res := []rune(strs[0])
			return string(res[0:length])
		}
	} else {
		return ""
	}
}

func main() {
	//fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	//fmt.Println(longestCommonPrefix([]string{"reflower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"aaa", "aa", "aaa"}))
}
