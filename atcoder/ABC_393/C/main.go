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
	m := ScanI()

	exist := make(map[string]bool)

	ans := 0

	for i := 0; i < m; i++ {
		u := ScanS()
		v := ScanS()
		if n == n {
		}
		// 小さい順にスワップ
		if v < u {
			u, v = v, u
		}
		key := u + "/" + v
		// 自己ループ
		if u == v {
			ans++
			continue
		}
		// 多重辺
		if _, ok := exist[key]; ok {
			ans++
			continue
		} else {
			exist[key] = true
		}
	}

	fmt.Println(ans)
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
