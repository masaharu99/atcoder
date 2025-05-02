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
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = ScanIArrayWithBlank(i + 1)
	}

	cur := 1
	for i := 1; i < n+1; i++ {
		if cur >= i {
			cur = a[cur-1][i-1]
		} else {
			cur = a[i-1][cur-1]
		}
		// fmt.Printf("%d回目：cur=%d\n", i, cur)
	}
	fmt.Println(cur)
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
