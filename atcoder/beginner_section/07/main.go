package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	
	d := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&d[i])
	}

	var ans int
	m := make(map[int]int)


	for i:=0; i<n; i++ {
		val := d[i]
		_, ok := m[val]
		if !ok {
			m[val] = 0
			ans++
		}
	}

	fmt.Println(ans)
}