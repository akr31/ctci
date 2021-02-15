package main

import "fmt"

func urlify(s string) string {
	var url string = ""
	var bufferFlag bool
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			url = string(s[i]) + url
			bufferFlag = true
		} else if bufferFlag {
			url = "%20" + url
		}
	}
	return url
}

func main() {
	fmt.Println(urlify("/test slug/example 1"))
	fmt.Println(urlify("abc   "))
	fmt.Println(urlify("a bc  "))
}
