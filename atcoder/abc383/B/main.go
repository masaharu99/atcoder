package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var (
		h, w, d int
		s [][]string
	)
	fmt.Scan(&h, &w, &d)
	for i:=0; i<h; i++ {
		var line string
		fmt.Scan(&line)
		s = append(s, strings.Split(line, ""))
	}

	max := 0
	
	for i1 :=0; i1 < h; i1++ {
		for j1 :=0; j1 < w; j1++ {
			if s[i1][j1] == "#" {
				continue
			}
			for i2 :=0; i2 < h; i2++ {
				for j2 :=0; j2 < w; j2++ {
					if s[i2][j2] == "#" {
						continue
					}
					num := count(s, i1, j1, i2, j2, d)
					if max < num {
						max = num
					}
				}
			}
		}
	}

	fmt.Println(max)
}

func count(s [][]string, i1 int, j1 int, i2 int, j2 int, d int) int {
	list := make(map[string]int)
	for i :=0; i < len(s); i++ {
		for j :=0; j < len(s[0]); j++ {
			dx1 := math.Abs(float64(i1-i))
			dy1 := math.Abs(float64(j1-j))
			dx2 := math.Abs(float64(i2-i))
			dy2 := math.Abs(float64(j2-j))
			dn := math.Min(dx1 + dy1, dx2 + dy2)
			if i1==0 && j1==0 && i2==0 && j2==4 {
			}

			key := strconv.Itoa(i) + strconv.Itoa(j)
			_, ok := list[key]
			if dn<=float64(d) && s[i][j] == "." && !ok {
				list[key] = 0
			}
		}
	}

	return len(list)
}