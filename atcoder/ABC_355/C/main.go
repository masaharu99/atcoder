package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/liyue201/gostl/ds/set"
	"github.com/liyue201/gostl/utils/comparator"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	n, t := ScanI(), ScanI()
	a := ScanIArrayWithBlank(t)

	ds1 := set.New(comparator.IntComparator, set.WithGoroutineSafe())
	ds2 := set.New(comparator.IntComparator, set.WithGoroutineSafe())
	for i := 1; i < n*n+1; i += n + 1 {
		ds1.Insert(i)
	}
	for i := n; i < 1+n*(n-1)+1; i += n - 1 {
		ds2.Insert(i)
	}

	ver, hor := make([]int, n+1), make([]int, n+1)
	dia1, dia2 := n, n
	for i := 0; i < n+1; i++ {
		ver[i] = n
		hor[i] = n
	}

	for i, v := range a {
		ni := v / n
		if v%n != 0 {
			ni++
		}
		ver[ni]--

		nj := v - n*(ni-1)
		hor[nj]--

		if ds1.Contains(v) {
			dia1--
		}
		if ds2.Contains(v) {
			dia2--
		}

		if ver[ni] == 0 || hor[nj] == 0 || dia1 == 0 || dia2 == 0 {
			fmt.Println(i + 1)
			return
		}
	}

	fmt.Println(-1)
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
