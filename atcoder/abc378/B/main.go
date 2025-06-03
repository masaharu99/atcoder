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
	waste := make([][]int, n)
	for i := 0; i < n; i++ {
		q := ScanI()
		r := ScanI()
		waste[i] = []int{q, r}
	}
	q := ScanI()
	questions := make([][]int, q)
	for i := 0; i < q; i++ {
		t := ScanI()
		d := ScanI()
		questions[i] = []int{t, d}
	}

	var ans int
	for _, question := range questions {
		tq, dq := question[0], question[1]
		qn, rn := waste[tq-1][0], waste[tq-1][1]
		res := dq%qn
		if res == rn {
			fmt.Println(dq)
		} else if res < rn {
			ans = dq + (rn-res)
			fmt.Println(ans)
		} else {
			ans = dq + qn - (res-rn)
			fmt.Println(ans)
		}
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