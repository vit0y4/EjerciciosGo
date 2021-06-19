package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	dic := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	if len(s) <= 15 && len(s) >= 1 {
		split := strings.Split(s, "")
		sum := 0

		if _, ok := dic[s]; ok {
			return dic[s]
		} else {
			count := 0
			for k := 0; k < len(split); k++ {
				if _, ok := dic[split[k]]; ok {
					if k+1 < len(split) {
						if dic[split[k]] < dic[split[k+1]] {
							count = 0
							sum -= dic[split[k]]
						} else {
							if dic[split[k]] == dic[split[k+1]] {
								count = count + 1
								if count > 3 {
									return -1
								}
							}
							count = 0
							sum += dic[split[k]]
						}
					} else {
						sum += dic[split[k]]
					}
				}
			}
			return sum
		}
	} else {
		return -1
	}
}

func main() {
	/* fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IIII"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("LVIIII"))
	fmt.Println(romanToInt("MMMCMXCIX"))
	fmt.Println(romanToInt("MMMM")) */
	fmt.Println(romanToInt("MDCCCLXXXIV"))
}
