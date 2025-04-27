package main

import (
	"fmt"
	"sort"
)


func main_bk() {
	n := ScanI()
	a := ScanIArrayWithBlank(n)
	b := ScanIArrayWithBlank(n-1)

	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})

	misMatch := false
	ans := 0

	for i := 0; i < n; i++ {
		if i == n-1 {
			if misMatch {
				if a[i] > b[i] {
					ans = -1
				}
			} else {
				ans = a[i]
			}
			break
		}

		if a[i] > b[i] {
			if misMatch {
				ans = -1
				break
			} else {
				misMatch = true
				ans = a[i]
				cp := make([]int, len(b)-i+1)
				copy(cp, b[i:])
				b = append(b[:i], n)
				b = append(b, cp...)
			}
		}
	}

	fmt.Println(ans)
}
