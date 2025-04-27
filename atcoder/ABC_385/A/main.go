package main

import (
	"fmt"
)

func main() {
	var a,b,c int
	fmt.Scan(&a,&b,&c)

	if (a == b) {
		if (a == c) {
			fmt.Println("Yes")
			return
		}
	}

	if (a < b) {
		if (b < c) {
			if (a+b == c) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		} else {
			if (a+c == b) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	} else if (a < c) {
		if (a+b == c) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else {
		if (c+b == a) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}