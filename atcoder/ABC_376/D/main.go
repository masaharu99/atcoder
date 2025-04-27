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
var n int
var m int
var sideMap map[int][]int
var visited []bool
var ans = int(2e6)
func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	n = ScanI()
	m = ScanI()
	sideMap = make(map[int][]int)
	for i := 0; i < m; i++ {
		a := ScanI() - 1
		b := ScanI() - 1
		if _, ok := sideMap[a]; ok {
			sideMap[a] = append(sideMap[a], b)
		} else {
			sideMap[a] = []int{b}
		}
	}
	
	d := make([]int, n)
	q := queue.New[int]()
	q.Push(0)

	for !q.Empty() {
		cur := q.Pop()
		for _, next := range sideMap[cur] {
			if next == 0 {
				fmt.Println(d[cur] + 1)
				return
			}
			if d[next] == 0 {
				d[next] = d[cur] + 1
				q.Push(next)
			}
		}
	}

	fmt.Println(-1)
}

func dfs(next, num int, isFirst bool) {
	num++
	if next == 1 && !isFirst {
		ans = int(math.Min(float64(ans), float64(num)))
		return
	}

	 // すでに訪問している場合（ただし、next が開始ノードの場合は許容する）
	 if visited[next] && !(next == 1 && !isFirst) {
		return
	}

	visited[next] = true
	for _, val := range sideMap[next] {
		dfs(val, num, false)
	}
	visited[next] = false

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