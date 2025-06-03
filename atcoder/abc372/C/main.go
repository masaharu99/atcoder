package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	_, q := ScanI(), ScanI()
	s := strings.Split(ScanS(), "")
	x, c := make([]int, q), make([]string, q)
	for i := 0; i < q; i++ {
		x[i] = ScanI()
		c[i] = ScanS()
	}

	cnt := strings.Count(strings.Join(s, ""), "ABC")

	for i := 0; i < q; i++ {
		xi, ci := x[i]-1, c[i]
		if judge(s, xi) && cnt != 0 {
			cnt--
		}
		s[xi] = ci
		if judge(s, xi) {
			cnt++
		}
		fmt.Println(cnt)
	}
}

func judge(s []string, x int) bool {
	back := len(s) - x - 1

	if 2 <= back && (s[x] == "A" && s[x+1] == "B" && s[x+2] == "C") {
		return true
	}
	if (1 <= x && 1 <= back) && (s[x-1] == "A" && s[x] == "B" && s[x+1] == "C") {
		return true
	}
	if 2 <= x && (s[x-2] == "A" && s[x-1] == "B" && s[x] == "C") {
		return true
	}

	return false
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
