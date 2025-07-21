package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/liyue201/gostl/ds/set"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	t := ScanI()

	for i := 0; i < t; i++ {
		n := ScanI()
		s := "0" + ScanS()

		ok := make([]bool, 1<<n)
		ok[0] = true

		for j := 0; j < 1<<n; j++ {
			if !ok[j] {
				continue
			}

			for k := 0; k < n; k++ {
				// すでに混ぜている
				if j&(1<<k) != 0 {
					continue
				}

				next := j | (1 << k)
				// fmt.Printf("debug: next=%d, s=%s\n", next, s)
				if s[next] == '0' {
					// fmt.Println("merge")
					ok[next] = true
				}
			}
		}

		// fmt.Println(ok)
		if ok[(1<<n)-1] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func dfs(cur string, next int, merged []bool, mgNum int, danger *set.Set[string]) bool {
	mgNum++
	nCur := cur[:next-1] + "1" + cur[next:]

	if danger.Contains(nCur) {
		return false
	}

	if mgNum == len(merged) {
		return true
	}
	merged[next-1] = true

	for i := 0; i < len(merged); i++ {
		if merged[i] {
			continue
		}
		if dfs(nCur, i+1, merged, mgNum, danger) {
			return true
		}
	}
	merged[next-1] = false

	return false
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

func ReverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
