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

	var ans [][]int
	var a []int

	dfs(a, n, m, &ans)

	fmt.Println(len(ans))
	for _, v := range ans {
		for _, v2 := range v {
			fmt.Printf("%d " , v2)
		}
		fmt.Println()
	}
}
	
func dfs(a []int, n, m int, ans *[][]int) {
	if (len(a) == n) {
		na := make([]int, len(a))
		copy(na, a)
		*ans = append(*ans, na)
		return
	}

	l := 1
	if len(a) > 0 {
		l = a[len(a)-1] + 10
	}
	a = append(a, l)

	for a[len(a)-1]+10*(n-len(a)) <= m {
		dfs(a, n, m, ans)
		a[len(a)-1]++
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