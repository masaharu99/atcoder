package main

import "fmt"

func main() {
	var a,b,c,x int
	fmt.Scan(&a,&b,&c,&x)

	var ans int
	for i:=0; i<=a; i++ {
		for j:=0; j<=b; j++ {
			for k:=0; k<=c; k++ {
				sum := 500*i + 100*j + 50*k
				if sum == x {
					ans++
				}
			}
		}
	}

	fmt.Println(ans)
}