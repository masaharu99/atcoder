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

func main() {
	h := ScanI()
	w := ScanI()
	k := ScanI()
	s := make([][]string, h)
	isPassed := make([][]bool, h)
	for i := 0; i < h; i++ {
		s[i] = strings.Split(ScanS(), "")
		isPassed[i] = make([]bool, w)
	}

	start := queue.New[[]int]()
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "." {
				start.Push([]int{i,j})
			}
		}
	}

	var ans int
	var dfs func(int, int, int)
	dfs = func (ni, nj, num int)  {
		if ni<0 || nj<0 || h<=ni || w<=nj {
			return
		}
		// すでに通っていた or 障害物
		if isPassed[ni][nj] || s[ni][nj]=="#" {
			return
		}
		// 目標回数に到達
		if num == k-1 {
			ans++
			return
		}
		isPassed[ni][nj] = true
		num++

		dfs(ni-1, nj, num)
		dfs(ni+1, nj, num)
		dfs(ni, nj-1, num)
		dfs(ni, nj+1, num)

		isPassed[ni][nj] = false
	}

	for !start.Empty() {
		ij := start.Pop()
		dfs(ij[0], ij[1], -1)
	}

	fmt.Println(ans)
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