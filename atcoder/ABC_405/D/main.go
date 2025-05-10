package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/liyue201/gostl/ds/queue"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

type Queue struct {
	i int
	j int
}

func main() {
	h, w := ScanI(), ScanI()
	s := make([][]string, h)
	q := queue.New[Queue]()
	for i := 0; i < h; i++ {
		in := ScanS()
		s[i] = strings.Split(in, "")
		if strings.Contains(in, "E") {
			for j := 0; j < w; j++ {
				if in[j] == 'E' {
					q.Push(Queue{i, j})
				}
			}
		}
	}

	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for !q.Empty() {
		t := q.Pop()
		for _, d := range directions {
			ni, nj := t.i+d[0], t.j+d[1]
			if ni < 0 || h <= ni || nj < 0 || w <= nj {
				continue
			}
			if s[ni][nj] == "." {
				q.Push(Queue{ni, nj})
				s[ni][nj] = getAllow(d[0], d[1])
			}
		}
	}

	for _, v := range s {
		fmt.Println(strings.Join(v, ""))
	}
}

func getAllow(di, dj int) string {
	if di == 0 && dj == 1 {
		return "<"
	} else if di == 0 && dj == -1 {
		return ">"
	} else if di == 1 && dj == 0 {
		return "^"
	} else {
		return "v"
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
