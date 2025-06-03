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
	n := ScanI()

	l1 := make([][]string, 3)
	for i := 0; i < 3; i++ {
		if i == 1 {
			l1[i] = []string{"#", ".", "#"}
		} else {
			l1[i] = []string{"#", "#", "#"}
		}
	}

	if n == 0 {
		fmt.Println("#")
		return
	} else if n == 1 {
		for i := 0; i < len(l1); i++ {
			for j := 0; j < len(l1); j++ {
				fmt.Print(l1[i][j])
			}
			fmt.Println()
		}
		return
	}

	s := f(l1, 1, n)

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			fmt.Print(s[i][j])
		}
		fmt.Println()
	}
}

func f(ps [][]string, l, gl int) [][]string {
	wh := len(ps) * 3
	s := make([][]string, wh)
	for i := 0; i < wh; i++ {
		s[i] = make([]string, wh)
		for j := 0; j < wh; j++ {
			start := len(ps)
			end := len(ps)*2 - 1
			if (start <= i && i <= end) && (start <= j && j <= end) {
				s[i][j] = "."
			} else {
				s[i][j] = ps[i%len(ps)][j%len(ps)]
			}
		}
	}

	l++
	if l == gl {
		return s
	}

	return f(s, l, gl)
}

// func f(n int, wl [][]bool) {
// 	cut := int(math.Pow(3, float64(n-1)))
// 	for i := cut + 1; i < cut*2+1; i++ {
// 		for j := cut + 1; j < cut*2+1; j++ {
// 			wl[i][j] = true
// 		}
// 	}
// }

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

type UnionFind struct {
	n    int
	root []int
}

func NewUnionFind(n int) UnionFind {
	root := make([]int, n+1)
	for i := 0; i < n; i++ {
		root[i] = -1
	}

	return UnionFind{n, root}
}

func (uf UnionFind) Leader(x int) int {
	if uf.root[x] < 0 {
		return x
	}

	uf.root[x] = uf.Leader(uf.root[x])
	return uf.root[x]
}

func (uf UnionFind) Merge(x, y int) int {
	xr := uf.Leader(x)
	yr := uf.Leader(y)
	if xr == yr {
		return xr
	}

	if -uf.root[xr] < -uf.root[yr] {
		xr, yr = yr, xr
	}
	uf.root[xr] += uf.root[yr]
	uf.root[yr] = xr

	return xr
}

func (uf UnionFind) Same(x, y int) bool {
	if uf.Leader(x) == uf.Leader(y) {
		return true
	}

	return false
}
