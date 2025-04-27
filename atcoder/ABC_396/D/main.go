package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/liyue201/gostl/ds/stack"
)

var sc = bufio.NewScanner(os.Stdin)
var (
	n   int
	m   int
	g   = make([][]hen, 11)
	th  = make([]bool, 11)
	min = -1
	s   = stack.New[int]()
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

type hen struct {
	end   int
	label int
}

func main() {
	n = ScanI()
	m = ScanI()
	for i := 0; i < m; i++ {
		a := ScanI()
		b := ScanI()
		l := ScanI()
		g[a] = append(g[a], hen{end: b, label: l})
		g[b] = append(g[b], hen{end: a, label: l})
	}

	dfs(1, 0)
	fmt.Println(min)
}

func dfs(next int, sum int) {
	ng := g[next]
	if len(ng) == 0 || th[next] == true {
		return
	}

	if next == n {
		if min == -1 {
			min = sum
		} else if sum < min {
			min = sum
		}
		return
	}

	th[next] = true
	s.Push(next)
	for _, val := range ng {
		dfs(val.end, sum^val.label)
	}
	th[next] = false
	s.Pop()
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
