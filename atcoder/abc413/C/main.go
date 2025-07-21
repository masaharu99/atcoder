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
	q := ScanI()
	queryList := make([][]int, q)

	for i := 0; i < q; i++ {
		t := ScanI()
		if t == 1 {
			queryList[i] = []int{t, ScanI(), ScanI()}
		} else {
			queryList[i] = []int{t, ScanI()}
		}
	}

	// 値、個数の順番
	vl := make([][]int, 0, q)
	head := 0

	for _, query := range queryList {
		if query[0] == 1 {
			vl = append(vl, []int{query[2], query[1]})
			continue
		}
		// fmt.Printf("debug: vl=%v, head=%d\n", vl, head)

		cut := query[1]
		ans := 0
		for cut > 0 {
			if vl[head][1] > cut {
				ans += vl[head][0] * cut
				vl[head][1] -= cut
				cut = 0
			} else {
				v := vl[head][0]
				n := vl[head][1]
				ans += v * n
				cut -= n
				head++
			}
		}
		fmt.Println(ans)
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

func ReverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
