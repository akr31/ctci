package main

import (
	"fmt"
	"reflect"
)

func getStringMap(s string) map[rune]int {
	var strMap map[rune]int = make(map[rune]int)
	for _, c := range s {
		_, ok := strMap[c]
		if ok {
			strMap[c]++
		} else {
			strMap[c] = 1
		}
	}
	return strMap
}
func checkPerms(s1, s2 string) bool {
	map1 := getStringMap(s1)
	map2 := getStringMap(s2)
	return reflect.DeepEqual(map1, map2)
}

func main() {
	fmt.Println(checkPerms("123", "312"))
	fmt.Println(checkPerms("122", "124"))
}
