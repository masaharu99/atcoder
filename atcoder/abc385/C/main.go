package main

import (
	"fmt"
)

type Count struct {
	n int
}

func main() {
	var n int
	fmt.Scan(&n)
	h := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&h[i])
	}

	ans := 0
	for i:=1; i<=n; i++ {
		for ini:=0; ini<i; ini++ {
			prev := 0
			count := 0
			for j:=ini; j<n; j+=i {
				if prev == h[j] {
					count++
				} else {
					count = 1
					prev = h[j]
				}
				if ans < count {
					ans = count
				}
			}
		}
	}

	fmt.Println(ans)
}