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

	a := 1
	ans := 0
	for i := 1; ; i++ {
		a *= 2
		if N < a {
			break
		}

		for j := 1; ; j += 2 {
			b := j * j
			if N < a*b {
				break
			}
			ans++
		}
	}

	fmt.Println(ans)
}

func getA(max int) []int {
	var a []int
	n := 1
	for i := 1; ; i++ {
		n *= 2
		if max < n {
			break
		}
		a = append(a, n)
	}

	return a
}

func getB(max int) []int {
	var b []int
	nMax := max / 2
	for i := 1; ; i++ {
		add := i * i
		if nMax < add {
			break
		}
		b = append(b, add)
	}

	return b
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
