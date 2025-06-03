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
	p := ScanIArrayWithBlank(n)
	np := make([]int, n)
	copy(np, p)

	sort.Slice(np, func(i, j int) bool {
		return np[i] > np[j]
	})

	prev := -1
	rank := make([]int, 101)
	cur := 0
	for i := 0; i < n; i++ {
		num := np[i]
		cur++

		if prev == num {
			continue
		}

		rank[num] = cur
		prev = num
	}

	for i := 0; i < n; i++ {
		n := p[i]
		fmt.Println(rank[n])
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
