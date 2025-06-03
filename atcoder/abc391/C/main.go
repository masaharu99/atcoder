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
	q := ScanI()
	query := make([][]int, q)
	for i := 0; i < q; i++ {
		f := ScanI()
		query[i] = []int{f}
		if f == 2 {
			continue
		}
		a := ScanIArrayWithBlank(2)
		query[i] = append(query[i], a...)
	}

	bird := make([]int, n)
	su := make([]int, n)
	for i := 0; i < n; i++{
		bird[i] = i
		su[i] = 1
	}

	ans := 0

	for _, v := range query {
		if v[0] == 2 {
			fmt.Println(ans)
		} else if v[0] == 1 {
			tg := v[1]-1
			after := v[2]-1
			current := bird[tg]
			bird[tg] = after
			su[current] -= 1
			su[after] += 1
			if su[current] == 1 {
				ans = int(math.Max(0, float64(ans-1)))
			}
			if su[after] == 2 {
				ans++
			}
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