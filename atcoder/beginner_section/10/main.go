package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	t := make([][]int, n)
	for i:=0; i<n; i++ {
		var tn, xn, yn int
		fmt.Scanf("%d %d %d", &tn, &xn, &yn)
		t[i] = []int{tn, xn, yn}
	}

	ct,cx,cy := 0,0,0
	for _,next := range t {
		for ; ct<next[0]; ct++ {
			switch {
			case cx < next[1]:
				cx++
			case cx > next[1]:
				cx--
			case cy < next[2]:
				cy++
			case cy > next[2]:
				cy--
			default:
				cx++
			}
		}
		
		if cx!=next[1] || cy!=next[2] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}