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
	a := ScanIArrayWithBlank(n)
	prev := make([]int, 1000001)
	ans := 10000010

	for i := 0; i < len(a); i++ {
		ni := i + 1
		val := a[i]
		if prev[val] != 0 {
			d := ni - prev[val] + 1
			// fmt.Println(prev[val], val, ans, d)
			ans = int(math.Min(float64(ans), float64(d)))
		}
		prev[val] = ni
	}

	if ans == 10000010 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
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
