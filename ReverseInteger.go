package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	///Solucion 1: menor tiempo de ejecucion
	///pero mas complicado
	/* var s string
	if x < 0 {
		s = strconv.Itoa(-x)
	} else {
		s = strconv.Itoa(x)
	}
	v := strings.SplitAfter(s, "")
	var nuevo string
	for i := len(v) - 1; i >= 0; i-- {
		nuevo += v[i]
	}
	numero, _ := strconv.Atoi(nuevo)
	if numero < math.MaxInt32 {
		if x < 0 {
			return -numero
		} else {
			return numero
		}
	} else {
		return 0
	} */
	/// solucion 2 menos complicado pero mayor tiempo de ejecucion
	sol := 0
	var n int
	if x < 0 {
		n = -x
	} else {
		n = x
	}
	for n > 0 {
		ultimo := n % 10
		sol = (sol * 10) + ultimo
		n = n / 10
	}
	if sol < math.MaxInt32 {
		if x < 0 {
			return -sol
		} else {
			return sol
		}
	} else {
		return 0
	}

}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(math.MinInt32 < (-321))
	fmt.Println(math.MaxInt32 > (321))
}
