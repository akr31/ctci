package main

import "fmt"

func isUnique(s string) bool {

	var strMap map[rune]bool = make(map[rune]bool)
	for _, c := range s {
		_, ok := strMap[c]
		if ok {
			return false
		}
		strMap[c] = true
	}
	return true
}

func main() {
	fmt.Println(isUnique("test"))
	fmt.Println(isUnique("ttest"))
}
