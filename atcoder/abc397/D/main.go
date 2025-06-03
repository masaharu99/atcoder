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
	nn := int64(n)

	for d := int64(1); d*d*d < nn; d++ {
		if nn%d != 0 {
			continue
		}

		k := f(3, 3*d, d*d-nn/d)
		if 0 < k {
			fmt.Println(k+d, k)
			return
		}
	}

	fmt.Println(-1)
}

func f(a int64, b int64, c int64) int64 {
	l := int64(0)
	r := int64(600000001)
	for 1 < r-l {
		// midがk（解）
		mid := (l + r) / 2
		if a*mid*mid+b*mid+c <= 0 {
			l = mid
		} else {
			r = mid
		}
	}

	if a*l*l+b*l+c == 0 {
		return l
	}

	return -1
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
