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

type Pair struct {
	x int
	y int
}

func main() {
	x, y := ScanI(), ScanI()
	s := set.New(comparator.StringComparator, set.WithGoroutineSafe())

	for i := 1; i < 7; i++ {
		for j := 1; j < 7; j++ {
			if x <= i+j {
				s.Insert(strconv.Itoa(i) + "," + strconv.Itoa(j))
			}

			d := j - i
			if j < i {
				d = i - j
			}
			if y <= d {
				s.Insert(strconv.Itoa(i) + "," + strconv.Itoa(j))
			}
		}
	}

	fmt.Println(float64(s.Size()) / 36)
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
