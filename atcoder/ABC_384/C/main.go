package main

import (
	"fmt"
	"sort"
)

func main() {
	var a,b,c,d,e int
	fmt.Scan(&a, &b, &c, &d, &e)
	problems := []struct {
		name string
		score int
	}{
		{"A", a},
		{"B", b},
		{"C", c},
		{"D", d},
		{"E", e},
	}

	type Users struct {
		name string
		score int
	}
	var users []Users
	for mask:=1; mask < (1<<len(problems)); mask++ {
		name := ""
		score := 0
		for i:=0; i<len(problems); i++ {
			if mask>>i&1 == 1 {
				name += string('A' + i)
				score += problems[i].score
			}
		}
		users = append(users, Users{name: name, score: score})
	}

	sort.Slice(users, func(i, j int) bool {
		if users[i].score != users[j].score {
			return users[i].score > users[j].score
		}
		return users[i].name < users[j].name
	})

	for _,user := range users {
		fmt.Println(user.name)
	}
}