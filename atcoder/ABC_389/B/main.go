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
	x := ScanI()

	for i := 0; true; i++ {
		val := factorial(i)
		if x == val {
			fmt.Println(i)
			break
		}
	}
}

func permutation(n int, k int) int {
	v := 1
	if 0 < k && k <= n {
			for i := 0; i < k; i++ {
					v *= (n - i)
			}
	} else if k > n {
			v = 0
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n - 1)
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