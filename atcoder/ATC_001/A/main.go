package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	h := ScanI()
	w := ScanI()
	c := make([][]string, h)
	for i := 0; i < h; i++ {
		c[i] = strings.Split(ScanS(), "")
	}

	solve(h, w, c)
}

func solve(h, w int, c[][]string) {
	grid := make([][]bool, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]bool, w)
	}
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var sx, sy int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if c[i][j] == "s" {
				sy = i
				sx = j
			}
		}
	}

	ans := "No"
	var dfs func(int, int)
	dfs = func (x, y int) {
		if x < 0 || x >= w || y < 0 || y >= h {
			return
		}
		if grid[y][x] {
			return
		}
		if c[y][x] == "#" {
			return
		}
		if c[y][x] == "g" {
			ans = "Yes"
			return
		}
	
		grid[y][x] = true

		for _, d := range directions {
			dfs(x+d[0], y+d[1])
		}
	}

	dfs(sx, sy)
	fmt.Println(ans)
}

func ScanI() int {
	sc.Scan()
	str, _ := strconv.Atoi(sc.Text())
	return str
}

func ScanS() string {
	sc.Scan()
	return sc.Text()
}

func ScanIArrayNotBlank() []int {
	str := strings.Split(ScanS(), "")
	arrI := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(str[i])
		arrI[i] = num
	}

	return arrI
}

func ScanIArrayWithBlank(n int) []int {
	arrI := make([]int, n)
	for i := 0; i < n; i++ {
		arrI[i] = ScanI()
	}

	return arrI
}