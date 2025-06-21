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
	n, q := ScanI(), ScanI()
	a := ScanIArrayWithBlank(q)

	// falseが白
	colors := make([]bool, n)
	cnt := 0

	for _, v := range a {
		c := colors[v-1]
		colors[v-1] = !colors[v-1]
		// fmt.Printf("debug: v=%d\n", v)

		// 黒塗り
		if !c {
			// 1マスのみ
			if n == 1 {
				cnt++
				fmt.Println(cnt)
				continue
			}
			// 左端
			if v-1 == 0 && !colors[v] {
				cnt++
				fmt.Println(cnt)
				continue
			}
			// 右端
			if v-1 == n-1 && !colors[n-2] {
				cnt++
				fmt.Println(cnt)
				continue
			}
			// 真ん中
			if 0 < v-1 && v-1 < n-1 {
				if !colors[v-2] && !colors[v] {
					cnt++
				} else if colors[v-2] && colors[v] {
					cnt--
				}
			}
			fmt.Println(cnt)
		} else {
			// 1マスのみ
			if n == 1 {
				cnt--
				fmt.Println(cnt)
				continue
			}
			// 左端
			if v-1 == 0 && !colors[v] {
				cnt--
				fmt.Println(cnt)
				continue
			}
			// 右端
			if v-1 == n-1 && !colors[n-2] {
				cnt--
				fmt.Println(cnt)
				continue
			}
			// 真ん中
			if 0 < v-1 && v-1 < n-1 {
				if !colors[v-2] && !colors[v] {
					cnt--
				} else if colors[v-2] && colors[v] {
					cnt++
				}
			}
			fmt.Println(cnt)
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
