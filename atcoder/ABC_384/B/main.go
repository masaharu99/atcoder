package main

import "fmt"

func main() {
	var (
		n, r int
	)
	fmt.Scan(&n, &r)

	dn := make([]int, n)
	an := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&dn[i], &an[i])
	}

	for i:=0; i<n; i++ {
		if dn[i] == 1 && 1600<=r && r<=2799 {
			r += an[i]
		} else if dn[i] == 2 && 1200<=r && r<=2399 {
			r += an[i]
		}
	}

	fmt.Println(r)
}