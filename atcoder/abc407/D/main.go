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
var (
	h, w int
	a    [][]int
	max  int
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	h, w = ScanI(), ScanI()
	a = make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = ScanIArrayWithBlank(w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			max = max ^ a[i][j]
		}
	}

	puted := make([][]bool, h)
	for i := 0; i < h; i++ {
		puted[i] = make([]bool, w)
	}

	dfs(puted, 0, 0)

	fmt.Println(max)
}

func dfs(puted [][]bool, ch, cw int) {
	if ch < 0 || h <= ch || cw < 0 || w <= cw {
		return
	}

	if cw == w-1 {
		dfs(puted, ch+1, 0)
	} else {
		dfs(puted, ch, cw+1)
	}

	if cw < w-2 && !puted[ch][cw] && !puted[ch][cw+1] {
		puted[ch][cw] = true
		puted[ch][cw+1] = true
		nmax := max ^ a[ch][cw]
		nmax = nmax ^ a[ch][cw+1]
		if max < nmax {
			max = nmax
		}
		fmt.Println(puted)

		if cw == w-1 {
			dfs(puted, ch+1, 0)
		} else {
			dfs(puted, ch, cw+1)
		}
		puted[ch][cw] = false
		puted[ch][cw+1] = false
	}

	if ch < h-2 && !puted[ch][cw] && !puted[ch+1][cw] {
		puted[ch][cw] = true
		puted[ch+1][cw] = true
		nmax := max ^ a[ch][cw]
		nmax = nmax ^ a[ch+1][cw]
		if max < nmax {
			max = nmax
		}
		fmt.Println(puted)

		if cw == w-1 {
			dfs(puted, ch+1, 0)
		} else {
			dfs(puted, ch, cw+1)
		}
		puted[ch][cw] = false
		puted[ch+1][cw] = false
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
