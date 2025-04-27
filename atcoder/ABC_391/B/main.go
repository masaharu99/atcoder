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
	n := ScanI()
	m := ScanI()
	s := make([]string, n)
	t := make([]string, m)
	for i := 0; i < n; i++ {
		s[i] = ScanS()
	}
	for i := 0; i < m; i++ {
		t[i] = ScanS()
	}

	for i := 0; i < n-m+1; i++ {
		for j:=0; j < n-m+1; j++ {
			if judge(i, j, s, t) {
				fmt.Println(i+1, j+1)
				return
			}
		}
	}
}

func judge(a, b int, s, t []string) bool {
	for i := a; i < a+len(t); i++ {
			if s[i][b:b+len(t)] != t[i-a] {
				return false
			}
	}

	return true
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