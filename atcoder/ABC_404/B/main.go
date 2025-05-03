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
	N := ScanI()
	S, T := make([][]string, N), make([][]string, N)
	for i := 0; i < N; i++ {
		S[i] = strings.Split(ScanS(), "")
	}
	for i := 0; i < N; i++ {
		T[i] = strings.Split(ScanS(), "")
	}

	a := rotation(S, N)
	b := rotation(a, N)
	c := rotation(b, N)

	n := solve(S, T)
	an := solve(a, T) + 1
	bn := solve(b, T) + 2
	cn := solve(c, T) + 3

	if an > n {
		an, n = n, an
	}
	if an > bn {
		an, bn = bn, an
	}
	if an > cn {
		an, cn = cn, an
	}

	fmt.Println(an)
}

func rotation(s [][]string, N int) [][]string {
	ns := make([][]string, N)
	for i := 0; i < N; i++ {
		ns[i] = make([]string, N)
	}
	// fmt.Println(len(ns))

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			// fmt.Println(j, N-1-i)
			// fmt.Println(ns[j])
			// fmt.Println(s[i][j])
			ns[j][N-1-i] = s[i][j]
		}
	}
	return ns
}

func solve(s, t [][]string) int {
	cnt := 0

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i][j] != t[i][j] {
				cnt++
			}
		}
	}

	return cnt
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
