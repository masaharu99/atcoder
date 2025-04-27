package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	
	a := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Slice(a, func(i, j int) bool {return a[i] > a[j]})

	var aliceSum int
	var bobSum int

	for i:=0; i<n; i++ {
		if i%2 == 0 {
			aliceSum += a[i]
		} else {
			bobSum += a[i]
		}
	}

	fmt.Println(aliceSum - bobSum)
}