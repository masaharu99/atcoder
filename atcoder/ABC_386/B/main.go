package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	ans := 0
	for {
		if s == "" {
			break
		}
		if strings.HasPrefix(s, "00") {
			s = s[2:]
		} else {
			s = s[1:]
		}
		ans++
	}
	fmt.Println(ans)
}