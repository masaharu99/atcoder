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
	s := make([][]string, h)
	for i := 0; i < h; i++ {
		s[i] = strings.Split(ScanS(), "")
	}

	type Queue struct {
		i, j, act, cnt int
	}
	q := queue.New[Queue]()
	kv := make(map[string]int)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "S" {
				q.Push(Queue{i, j, -1, 0})
			}
		}
	}

	type Directions struct {
		i, j, act int
	}
	directions := []Directions {
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 1},
		{0, 1, 1},
	}

	for !q.Empty() {
		cq := q.Pop()

		// startの時
		if cq.act == -1 {
			for _,direction := range directions {
				ni := cq.i + direction.i
				nj := cq.j + direction.j
				// 範囲外の時 or 障害物の時
				if ni < 0 || nj < 0 || h <= ni || w <= nj || s[ni][nj] == "#" {
					continue
				}
				// すでに通っているとき
				key := strconv.Itoa(ni) + "/" + strconv.Itoa(nj) + strconv.Itoa(cq.act)
				_, ok := kv[key]
				if ok {
					continue
				}
				if s[ni][nj] == "." {
					q.Push(Queue{ni, nj, direction.act, cq.cnt+1})
					kv[key] = 0
				}
				if s[ni][nj] == "G" {
					fmt.Println(cq.cnt+1)
					return
				}
			}
		}

		// 前が縦移動の時
		if cq.act == 0 {
			for i := 2; i < 4; i++ {
				ni := cq.i + directions[i].i
				nj := cq.j + directions[i].j
				// 範囲外の時 or 障害物の時
				if ni < 0 || nj < 0 || h <= ni || w <= nj || s[ni][nj] == "#" {
					continue
				}
				// すでに通っているとき
				key := strconv.Itoa(ni) + "/" + strconv.Itoa(nj) + strconv.Itoa(cq.act)
				_, ok := kv[key]
				if ok {
					continue
				}
				if s[ni][nj] == "." {
					q.Push(Queue{ni, nj, directions[i].act, cq.cnt+1})
					kv[key] = 0
				}
				if s[ni][nj] == "G" {
					fmt.Println(cq.cnt+1)
					return
				}
			}
		}

		// 前が横移動の時
		if cq.act == 1 {
			for i := 0; i < 2; i++ {
				ni := cq.i + directions[i].i
				nj := cq.j + directions[i].j
				// 範囲外の時 or 障害物の時
				if ni < 0 || nj < 0 || h <= ni || w <= nj || s[ni][nj] == "#" {
					continue
				}
				// すでに通っているとき
				key := strconv.Itoa(ni) + "/" + strconv.Itoa(nj) + strconv.Itoa(cq.act)
				_, ok := kv[key]
				if ok {
					continue
				}
				if s[ni][nj] == "." {
					q.Push(Queue{ni, nj, directions[i].act, cq.cnt+1})
					kv[key] = 0
				}
				if s[ni][nj] == "G" {
					fmt.Println(cq.cnt+1)
					return
				}
			}
		}
	}
	fmt.Println(-1)
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