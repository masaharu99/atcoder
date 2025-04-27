package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		n int
		aInt []int
	)
	fmt.Scan(&n)

	aStr := make([]string, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&aStr[i])
	}

	for i:=0; i<n; i++ {
		num, _ := strconv.Atoi(aStr[i])
		aInt = append(aInt, num)
	}

	var ans int
	for {
		for i:=0; i<n; i++ {
			if aInt[i]%2 == 0 {
				aInt[i] = aInt[i] / 2
			} else {
				fmt.Println(ans)
				return
			}
		}
		ans++
	}
}