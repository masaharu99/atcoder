package main

import (
	"fmt"
)

func main() {
	var s string
	divide := []string{"dream", "dreamer", "erase", "eraser"}
	fmt.Scan(&s)
	
	rs := reverse(s)
	rd := make([]string, 4)
	for i:=0; i<4; i++ {
		rd[i] = reverse(divide[i])
	}

	for i:=0; i<len(rs); {
		var canContinue bool
		for _,val := range rd {
			lenght := len(val)
			if len(rs[i:]) < lenght {
				continue
			}

			if val == rs[i:i+lenght] {
				i += lenght
				canContinue = true
			}
		}

		if !canContinue {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}