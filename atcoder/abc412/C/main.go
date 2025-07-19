package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	t := ScanI()
	sl := make([][]int, t)
	first := make([]int, t)
	last := make([]int, t)

	for i := 0; i < t; i++ {
		n := ScanI()
		tmp := ScanIArrayWithBlank(n)
		first[i] = tmp[0]
		last[i] = tmp[n-1]
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		sl[i] = tmp
	}

	for i, s := range sl {
		if first[i]*2 >= last[i] {
			fmt.Println(2)
			continue
		}
		cur := first[i]
		ans := 2
		for {
			v := binarySearch(s, cur)
			if v == cur {
				fmt.Println(-1)
				break
			}

			ans++
			if v*2 >= last[i] {
				fmt.Println(ans)
				break
			}
			cur = v
		}
	}
}

func binarySearch(s []int, cur int) int {
	l, r := 0, len(s)-1
	for {
		mid := (l + r) / 2
		if r-l == 1 {
			return s[l]
		}

		if s[mid] <= cur*2 {
			l = mid
		} else {
			r = mid
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

func ReverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
