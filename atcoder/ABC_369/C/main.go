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
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = ScanI()
	}

	d := math.MaxInt
	n := 1
	ans := 1
	for i := 0; i < N-1; i++ {
		if d == math.MaxInt {
			d = A[i] - A[i+1]
			n++
			ans += n
			// fmt.Printf("%d回目：n=%d, ans=%d\n", i+1, n, d)
			continue
		}

		if d == A[i]-A[i+1] {
			n++
			ans += n
		} else {
			d = A[i] - A[i+1]
			n = 2
			ans += n
		}
		// fmt.Printf("%d回目：n=%d, ans=%d\n", i+1, n, d)
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
