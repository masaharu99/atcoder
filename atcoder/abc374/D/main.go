package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var (
	n  int
	wt int
	mt int
	f  [][]int
	s  [][]int
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	n = ScanI()
	wt = ScanI()
	mt = ScanI()
	f = make([][]int, n)
	s = make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = []int{ScanI(), ScanI()}
		s[i] = []int{ScanI(), ScanI()}
	}
}

func dfs()

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
