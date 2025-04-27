package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/liyue201/gostl/ds/set"
	"github.com/liyue201/gostl/utils/comparator"
)

var sc = bufio.NewScanner(os.Stdin)
func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	n, m, sx, sy := ScanI(), ScanI(), ScanI(), ScanI()
	x, y := make([]int, n), make([]int, n)
	d, c := make([]string, m), make([]int, m)
	// yで検索して、x軸の家を管理
	xs := make(map[int]*set.Set[int])
	// xで検索して、y軸の家を管理
	ys := make(map[int]*set.Set[int])
	for i := 0; i < n; i++ {
		x[i], y[i] = ScanI(), ScanI()
		if _,exist := xs[y[i]]; !exist {
			xs[y[i]] = set.New[int](comparator.IntComparator, set.WithGoroutineSafe())
		}
		if _,exist := ys[x[i]]; !exist {
			ys[x[i]] = set.New[int](comparator.IntComparator, set.WithGoroutineSafe())
		}
		xs[y[i]].Insert(x[i])
		ys[x[i]].Insert(y[i])
	}
	for i := 0; i < m; i++ {
		d[i], c[i] = ScanS(), ScanI()
	}

	ans := 0

	f := func (start, end, key int, search, other map[int]*set.Set[int]) {
		if (start > end) {
			start, end = end, start
		}
		// 家が存在しない
		st, ok := search[key]
		if !ok {
			return
		}

		for {
			it := st.LowerBound(start)
			if !it.IsValid() {
				break
			}
			if end < it.Value() {
				break
			}
			st.Erase(it.Value())
			other[it.Value()].Erase(key)
			ans++
		}
	}

	for i := 0; i < m; i++ {
		cd, cc := d[i], c[i]
		nx, ny := sx, sy
		switch cd {
		case "U":
			ny += cc
		case "D":
			ny -= cc
		case "R":
			nx += cc
		case "L":
			nx -= cc
		}
		
		if sy == ny {
			f(sx, nx, sy, xs, ys)
		} else {
			f(sy, ny, sx, ys, xs)
		}
		sx, sy = nx, ny
	}

	fmt.Println(sx, sy, ans)
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