package main

import (
	"fmt"
	"strings"
)

/* func isValid(s string) bool {
	sol := strings.Split(s, "")
	pila := make([]string, 0)
	for i := 0; i < len(sol); i++ {
		switch string(sol[i]) {
		case "(":
			pila = append(pila, sol[i])
		case "[":
			pila = append(pila, sol[i])
		case "{":
			pila = append(pila, sol[i])
		case ")":
			if len(pila) > 0 {
				if pila[len(pila)-1] == "(" {
					if len(pila)-1 == 0 {
						pila = nil
					} else {
						pila = pila[:len(pila)-1]
					}
				} else {
					return false
				}
			} else {
				return false
			}
		case "]":
			if len(pila) > 0 {
				if pila[len(pila)-1] == "[" {
					if len(pila)-1 == 0 {
						pila = nil
					} else {
						pila = pila[:len(pila)-1]
					}
				} else {
					return false
				}
			} else {
				return false
			}
		case "}":
			if len(pila) > 0 {
				if pila[len(pila)-1] == "{" {
					if len(pila)-1 == 0 {
						pila = nil
					} else {
						pila = pila[:len(pila)-1]
					}
				} else {
					return false
				}
			} else {
				return false
			}
		default:
			return false
		}
	}
	if len(pila) == 0 {
		return true
	} else {
		return false
	}
}*/

/* func isValid(s string) bool {
	sol := strings.Split(s, "")
	pila := make([]string, 0)
	for i := 0; i < len(sol); i++ {
		if sol[i] == "(" || sol[i] == "[" || sol[i] == "{" {
			pila = append(pila, sol[i])
		} else if len(pila) > 0 {
			switch string(sol[i]) {
			case ")":
				if pila[len(pila)-1] == "(" {
					if len(pila)-1 == 0 {
						pila = nil
					} else {
						pila = pila[:len(pila)-1]
					}
				} else {
					return false
				}
			case "]":
				if pila[len(pila)-1] == "[" {
					if len(pila)-1 == 0 {
						pila = nil
					} else {
						pila = pila[:len(pila)-1]
					}
				} else {
					return false
				}
			case "}":
				if pila[len(pila)-1] == "{" {
					if len(pila)-1 == 0 {
						pila = nil
					} else {
						pila = pila[:len(pila)-1]
					}
				} else {
					return false
				}
			default:
				return false
			}
		} else {
			return false
		}
	}
	if len(pila) == 0 {
		return true
	} else {
		return false
	}
} */

func isValid(s string) bool {
	sol := strings.Split(s, "")
	pila := make([]string, 0)
	cierres := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	for i := 0; i < len(sol); i++ {
		if sol[i] == "(" || sol[i] == "[" || sol[i] == "{" {
			pila = append(pila, sol[i])
		} else if _, ok := cierres[sol[i]]; ok {
			if len(pila) > 0 {
				if pila[len(pila)-1] == cierres[sol[i]] {
					pila = pila[:len(pila)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}
	if len(pila) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	//fmt.Println(isValid("]"))
	fmt.Println(isValid("(])"))
	//fmt.Println(isValid("()[]{}"))
	//fmt.Println(isValid("([)]"))
	//fmt.Println(isValid("{[]}"))
}
