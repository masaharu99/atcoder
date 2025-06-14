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
var ans = math.MaxInt64

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

type Node struct {
	to int
	h  int
}

func main() {
	n, m := ScanI(), ScanI()
	a, b, w := make([]int, m), make([]int, m), make([]int, m)
	node := make([][]Node, n+1)
	for i := 0; i < m; i++ {
		a[i] = ScanI()
		b[i] = ScanI()
		w[i] = ScanI()
		node[a[i]] = append(node[a[i]], Node{b[i], w[i]})
	}

	visited := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		visited[i] = make([]bool, n+1)
	}

	bfs(0, 1, n, 0, visited, node)

	if ans == math.MaxInt64 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

func bfs(before, now, n, h int, visited [][]bool, node [][]Node) {
	if now == n && h < ans {
		ans = h
	}
	visited[now][before] = true

	for _, v := range node[now] {
		if !visited[v.to][now] {
			h = h ^ v.h
			bfs(now, v.to, n, h, visited, node)
		}
	}

	visited[now][before] = false
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
