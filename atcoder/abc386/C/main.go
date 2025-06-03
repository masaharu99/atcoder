package main

import "fmt"

func main() {
	var (
		k int
		s string
		t string
	)
	fmt.Scan(&k, &s, &t)

	if s == t {
		fmt.Println("Yes")
	}

	if len(s)-len(t) == -1 {
		// 文字数が少ない
		for i:=0; i<len(s); i++ {
			if s[i] != t[i] {
				s = s[:i] + string(t[i]) + s[i:]
				if s == t {
					fmt.Println("Yes")
				} else {
					fmt.Println("No")
				}
				return
			}
			if i == len(s)-1 {
				fmt.Println("Yes")
			}
		}
	// 文字数が多い
	} else if len(s)-len(t) == 1 {
		for i:=0; i<len(s); i++ {
			if i == len(s)-1 {
				fmt.Println("Yes")
				return
			}
			if s[i] != t[i] {
				s = s[:i] + s[i+1:]
				if s == t {
					fmt.Println("Yes")
				} else {
					fmt.Println("No")
				}
				return
			}
		}
	// 文字数が同じ
	} else if len(s)-len(t) == 0{
		for i:=0; i<len(s); i++ {
			if s[i] != t[i] {
				s = s[:i] + string(t[i]) + s[i+1:]
				if s == t {
					fmt.Println("Yes")
				} else {
					fmt.Println("No")
				}
				return
			}
		}
	} else {
		fmt.Println("No")
	}
}