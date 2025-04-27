package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main_bk() {
	var a,b,c,d,e int
	fmt.Scan(&a, &b, &c, &d, &e)

	names := []string {
		"A","B","C","D","E",
		"AB","AC","AD","AE","BC","BD","BE","CD","CE","DE",
		"ABC","ABD","ABE","ACD","ACE","ADE","BCD","BCE","BDE","CDE",
		"ABCD","ABCE","ABDE","ACDE","BCDE",
		"ABCDE",
	}
	user := make([][]string, 31)

	for i:=0; i<31; i++ {
		sum := 0
		name := names[i]
		if strings.Contains(name, "A") {
			sum += a
		}
		if strings.Contains(name, "B") {
			sum += b
		}
		if strings.Contains(name, "C") {
			sum += c
		}
		if strings.Contains(name, "D") {
			sum += d
		}
		if strings.Contains(name, "E") {
			sum += e
		}

		user[i] = []string{names[i], strconv.Itoa(sum)}
	}

	sort.Slice(user, func(i, j int) bool {
		scoreI,_ := strconv.Atoi(user[i][1])
		scoreJ,_ := strconv.Atoi(user[j][1])

		if scoreI != scoreJ {
			return scoreI > scoreJ
		}
		return user[i][0] < user[j][0]
	})
	
	for _,val := range user {
		fmt.Println(val[0])
	}
}