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
	s := make([]string, n)
	maxLen := 0
	for i := 0; i < n; i++ {
		s[i] = ScanS()
		if maxLen < len(s[i]) {
			maxLen = len(s[i])
		}
	}
	// fmt.Println()

	ans := make([][]string, maxLen)
	for i := 0; i < maxLen; i++ {
		ans[i] = make([]string, n)
		for j := 0; j < n; j++ {
			if len(s[n-j-1]) <= i {
				ans[i][j] = "*"
			} else {
				ans[i][j] = string(s[n-j-1][i])
			}
		}
	}

	for i := 0; i < maxLen; i++ {
		for j := len(ans[i]) - 1; -1 < j; j-- {
			if ans[i][j] == "*" {
				ans[i][j] = ""
			} else {
				break
			}
		}
	}

	for i := 0; i < maxLen; i++ {
		fmt.Println(strings.Join(ans[i], ""))
	}
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
