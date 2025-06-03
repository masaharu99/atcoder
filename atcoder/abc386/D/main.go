package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type paint struct {
	x int
	y int
	c string
}

func main() {
	var (
		nm []string
		cn []paint
	)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	nm = strings.Split(sc.Text(), " ")
	m,_ := strconv.Atoi(nm[1])

	for i:=0; i<m; i++ {
		sc.Scan()
		str := strings.Split(sc.Text(), " ")
		x,_ := strconv.Atoi(str[0])
		y,_ := strconv.Atoi(str[1])
		cn = append(cn, paint{x,y,str[2]})
	}

	sort.Slice(cn, func(i, j int) bool {
		if cn[i].x != cn[j].x {
			return cn[i].x < cn[j].x
		}
		return cn[i].y < cn[j].y
	})

	minY := int(math.Pow(10, 9))
	ans := "Yes"
	for _,val := range cn {
		if val.c == "W" {
			minY = val.y
		} else if minY <= val.y {
			ans = "No"
		}
	}

	fmt.Println(ans)
}
