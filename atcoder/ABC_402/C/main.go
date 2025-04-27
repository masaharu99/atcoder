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
	N, M := ScanI(), ScanI()
	K := make([]int, M)
	A := make([][]int, M)
	for i := 0; i < M; i++ {
		k := ScanI()
		K[i] = k
		A[i] = make([]int, k)
		for j := 0; j < k; j++ {
			A[i][j] = ScanI()
		}
	}
	B := make([]int, N+1)
	for i := 0; i < N; i++ {
		n := ScanI()
		B[n] = i + 1
	}

	ans := make([]int, N+1)
	for i := 0; i < M; i++ {
		max := 0

		for _, v := range A[i] {
			d := B[v]
			if max < d {
				max = d
			}
		}

		ans[max]++
	}

	c := 0
	for i := 1; i < len(ans); i++ {
		c += ans[i]
		fmt.Println(c)
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
