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
	n, k := ScanI(), ScanI()
	a := make([]int, n+1)
	if n < k {
		fmt.Println(1)
		return
	}

	for i := 0; i < k; i++ {
		a[i] = 1
	}

	sum := k
	for i := k; i < len(a); i++ {
		a[i] = sum % 1000000000
		sum -= a[i-k]
		sum += a[i]
		// fmt.Println(i, sum)
	}
	// fmt.Println(a)

	fmt.Println(a[n])
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
