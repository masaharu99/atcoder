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

// キューをライブラリ使用

var sc = bufio.NewScanner(os.Stdin)
func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	h := ScanI()
	w := ScanI()
	d := ScanI()
	var s [][]string
	for i := 0; i < h; i++ {
		line := strings.Split(ScanS(), "")
		s = append(s, line)
	}

	type node struct {
		i, j int
	}

	dist := make([][]int, h)
	queue := queue.New[node]()
	// queue := NewQueue(int(math.Pow(10, 6)))

	for i := 0; i < h; i++ {
		dist[i] = make([]int, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "H" {
				dist[i][j] = 0
				queue.Push(node{i,j})
			} else {
				dist[i][j] = -1
			}
		}
	}

	type Ck struct {
		h int
		w int
	}
	ck := []Ck{{-1,0}, {1,0}, {0,-1}, {0,1}}

	for !queue.Empty() {
		f := queue.Pop()
		dep := dist[f.i][f.j]
		for _, val := range ck {
			x := f.j + val.w
			y := f.i + val.h
			if x < 0 || y < 0 || w <= x || h <= y {
				continue
			}
			if s[y][x] == "." && dist[y][x] == -1 {
				dist[y][x] = dep + 1
				queue.Push(node{y, x})
			}
		}
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if dist[i][j] != -1 && dist[i][j] <= d {
				ans++
			}
		}
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