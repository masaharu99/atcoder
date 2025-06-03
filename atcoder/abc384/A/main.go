package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		n int
		a,b string
		s string
	)
	fmt.Scan(&n, &a, &b)
	fmt.Scan(&s)

	ans := make([]string, n)
	for _,val := range s {
		if val != rune(a[0]) {
			ans = append(ans, string(b))
		} else {
			ans = append(ans, string(a))
		}
	}

	fmt.Println(strings.Join(ans, ""))
}