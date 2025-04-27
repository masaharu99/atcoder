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
	n := ScanI()
	a := ScanIArrayWithBlank(n)
	b := ScanIArrayWithBlank(n-1)

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	judge := func(wj int) bool {
		nb := make([]int, n-1)
		copy(nb, b)
		nb = append(nb, wj)
		sort.Slice(nb, func(i, j int) bool {
			return nb[i] < nb[j]
		})

		for i := 0; i < n; i++ {
			if a[i] > nb[i] {
				return false
			}
		}
		return true
	}

	INF := 1001001001
	ac , wa := INF, 0
	for ac-wa > 1 {
		wj := (ac+wa) / 2
		if judge(wj) {
			ac = wj
		} else {
			wa = wj
		}
	}
	
	if ac == INF {
		ac = -1
	}
	fmt.Println(ac)
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