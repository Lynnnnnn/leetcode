package main

import (
	"fmt"
	"strings"
)

func strWithout3a3b(A int, B int) string {
	result := ""

	for {
		if A-B >= 3 {
			result += "aab"
			A -= 2
			B -= 1
		} else if A-B < 3 && A-B >= 0 {
			result += strings.Repeat("ab", B)
			result += strings.Repeat("a", A-B)
			break
		} else if A-B < 0 && A-B > -3 {
			result += strings.Repeat("ba", A)
			result += strings.Repeat("b", B-A)
			break
		} else if A-B <= -3 {
			result += "bba"
			A -= 1
			B -= 2
		}
	}

	return result
}

func main() {
	fmt.Printf(strWithout3a3b(24, 11))
}
