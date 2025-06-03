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

type Pair struct {
	cmp string
	num int
}

func main() {
	n := ScanI()
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = ScanI()
	}

	arr := make([]string, n-1)
	for i := 0; i < n-1; i++ {
		if p[i] < p[i+1] {
			arr[i] = "<"
		} else {
			arr[i] = ">"
		}
	}

	var cmpl []Pair
	for i, v := range arr {
		if i == 0 {
			cmpl = append(cmpl, Pair{v, 1})
			continue
		}
		prev := len(cmpl) - 1
		if cmpl[prev].cmp == v {
			cmpl[prev].num++
		} else {
			cmpl = append(cmpl, Pair{v, 1})
		}
	}

	ans := 0
	for i := 1; i < len(cmpl)-1; i++ {
		if cmpl[i].cmp != ">" {
			continue
		}
		if cmpl[i-1].cmp != "<" || cmpl[i+1].cmp != "<" {
			continue
		}
		ans += cmpl[i-1].num * cmpl[i+1].num
	}

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
