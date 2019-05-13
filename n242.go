package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := map[int]int{}
	tMap := map[int]int{}

	len := len(s)

	for i := 0; i < len; i++ {
		sMap[int(s[i])] += 1
		tMap[int(t[i])] += 1
	}

	for i := 0; i < len; i++ {
		if sMap[int(s[i])] != tMap[int(s[i])] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Print(isAnagram("abcbb", "abccc"))
}
