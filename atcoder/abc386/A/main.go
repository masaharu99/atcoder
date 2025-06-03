package main

import (
	"fmt"
)

func main() {
	s := make([]int, 4)
	for i:=0; i<4; i++ {
		fmt.Scan(&s[i])
	}

	counts := make(map[int]int)

	for i:=0; i<4; i++ {
		_,ok := counts[s[i]]
		if ok {
			counts[s[i]]++
		} else {
			counts[s[i]] = 1
		}
	}

	if len(counts) == 2 {
		for _,val := range counts {
			if val == 0 {
				fmt.Println("No")
				return
			}
		}
	} else {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}