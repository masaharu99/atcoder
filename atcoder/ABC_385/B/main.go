package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var H,W,X,Y int
	var S [][]string
	var T string
	fmt.Scan(&H,&W,&X,&Y)

	sc := bufio.NewScanner(os.Stdin)
	for i:=0; i<H; i++ {
		sc.Scan()
		S = append(S, strings.Split(sc.Text(), ""))
	}

	fmt.Scanf("%s", &T)

	cX := X-1
	cY := Y-1
	history := make(map[string]int)
	count := 0

	for _, c := range strings.Split(T, "") {
		if (c == "U") {
			if (S[cX-1][cY] == "#") {
				continue
			} else if (S[cX-1][cY] == ".") {
				cX = cX-1
			} else {
				cX = cX-1
				_, ok := history[getKey(cX, cY)]
				if (ok == false) {
					history[getKey(cX, cY)] = 0
					count++
				}
			}
		} else if (c == "D") {
			if (S[cX+1][cY] == "#") {
				continue
			} else if (S[cX+1][cY] == ".") {
				cX = cX + 1
			} else {
				cX = cX + 1
				_, ok := history[getKey(cX, cY)]
				if (ok == false) {
					history[getKey(cX, cY)] = 0
					count++
				}
			}
		} else if (c == "L") {
			if (S[cX][cY-1] == "#") {
				continue
			} else if (S[cX][cY-1] == ".") {
				cY = cY - 1
			} else {
				cY = cY - 1
				_, ok := history[getKey(cX, cY)]
				if (ok == false) {
					history[getKey(cX, cY)] = 0
					count++
				}
			}
		} else {
			if (S[cX][cY+1] == "#") {
				continue
			} else if (S[cX][cY+1] == ".") {
				cY = cY + 1
			} else {
				cY = cY + 1
				_, ok := history[getKey(cX, cY)]
				if (ok == false) {
					history[getKey(cX, cY)] = 0
					count++
				}
			}
		}
	}

	fmt.Println(cX+1, cY+1, count)
}

func getKey(x, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}