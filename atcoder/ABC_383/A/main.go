package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	tn := make([]int, n)
	vn := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&tn[i], &vn[i])
	}

	ans := 0
	prevt := 0
	for i:=0; i<n; i++{
		ans =  ans - (tn[i] - prevt)
		if ans < 0 {
			ans = 0
		}

		ans += vn[i]
		prevt = tn[i]
	}

	fmt.Println(ans)
}