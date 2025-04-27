package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	na := make([]int, n)
	copy(na, a)
	// arr := make([][]int, n)

	// for i := 0; i < n; i++ {
	// 	arr[i] = []int{a[i], i + 1}
	// }

	if n == 1 {
		fmt.Println(1)
		return
	}

	sort.Slice(na, func(i, j int) bool {
		return na[i] > na[j]
	})

	prev := -1
	cnt := 0
	max := -1

	for i, val := range na {
		if prev == val {
			cnt++
			continue
		}

		if i == len(na)-1 {
			max = val
		}

		if cnt == 1 {
			max = prev
			break
		} else {
			cnt = 1
			prev = val
		}
	}

	if max == -1 {
		fmt.Println(-1)
		return
	}

	for i, val := range a {
		if val == max {
			fmt.Println(i + 1)
			break
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
