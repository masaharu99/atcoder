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
	m := ScanI()

	n := 1
	l := []int{1}
	for i := 1; ; i++ {
		n *= 3
		if m < n {
			break
		}
		l = append(l, n)
	}

	ans := make([]int, 0, 20)
	for i := len(l) - 1; -1 < i; i-- {
		if m == 0 {
			break
		}
		for {
			if m < l[i] {
				break
			}
			// fmt.Printf("debug: m=%d, i=%d , l[i]=%d\n", m, i, l[i])
			m -= l[i]
			ans = append(ans, i)
		}
	}

	fmt.Println(len(ans))
	for _, v := range ans {
		fmt.Print(v, " ")
	}
	fmt.Println()
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
