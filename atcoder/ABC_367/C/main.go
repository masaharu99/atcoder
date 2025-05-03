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
var N, K int
var R []int
var ans = make([][]int, 0, int(math.Pow(5, 8)))

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	N, K = ScanI(), ScanI()
	R = make([]int, N)
	for i := 0; i < N; i++ {
		R[i] = ScanI()
	}

	add := make([]int, N)
	dfs(0, add, 0)

	for _, vl := range ans {
		for _, v := range vl {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

func dfs(i int, add []int, sum int) {
	if i == N {
		if sum%K == 0 {
			nAdd := make([]int, len(add))
			copy(nAdd, add)
			ans = append(ans, nAdd)
		}
		return
	}

	for j := 1; j < R[i]+1; j++ {
		add[i] = j
		sum += j
		dfs(i+1, add, sum)
		sum -= j
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
