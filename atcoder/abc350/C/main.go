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

type pair struct {
	i, j int
}

func main() {
	n := ScanI()
	a := make([]int, n)
	ngPosition := set.New(comparator.IntComparator, set.WithGoroutineSafe())
	numToPosition := make([]int, n+1)
	for i := 0; i < n; i++ {
		n := ScanI()
		a[i] = n
		numToPosition[n] = i
		if i+1 != n {
			ngPosition.Insert(i)
		}
	}

	cnt := 0
	arr := make([]pair, 0)
	for ngPosition.Size() != 0 {
		fp := ngPosition.First().Value()
		ngPosition.Erase(fp)
		cp := numToPosition[fp+1]

		arr = append(arr, pair{fp + 1, cp + 1})
		a[fp], a[cp] = a[cp], a[fp]
		cnt++

		if cp+1 == a[cp] {
			ngPosition.Erase(cp)
		}
	}

	fmt.Println(cnt)
	for _, v := range arr {
		fmt.Println(v.i, v.j)
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
