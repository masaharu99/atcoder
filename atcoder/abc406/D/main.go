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
	h, w, n := ScanI(), ScanI(), ScanI()
	xn := make([]*set.Set[int], h+1)
	yn := make([]*set.Set[int], w+1)
	for i := 0; i < h+1; i++ {
		xn[i] = set.New(comparator.IntComparator, set.WithGoroutineSafe())
	}
	for i := 0; i < w+1; i++ {
		yn[i] = set.New(comparator.IntComparator, set.WithGoroutineSafe())
	}

	for i := 0; i < n; i++ {
		x, y := ScanI(), ScanI()
		xn[x].Insert(y)
		yn[y].Insert(x)
	}
	q := ScanI()
	query := make([][]int, q)
	for i := 0; i < q; i++ {
		query[i] = []int{ScanI(), ScanI()}
	}

	for _, v := range query {
		q1 := v[0]
		q2 := v[1]
		if q1 == 1 {
			fmt.Println(xn[q2].Size())
			for xn[q2].Size() != 0 {
				y := xn[q2].First().Value()
				xn[q2].Erase(y)
				yn[y].Erase(q2)
			}
		} else {
			fmt.Println(yn[q2].Size())
			for yn[q2].Size() != 0 {
				x := yn[q2].First().Value()
				yn[q2].Erase(x)
				xn[x].Erase(q2)
			}
		}
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
