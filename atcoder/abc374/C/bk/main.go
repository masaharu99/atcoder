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
var n int
var k []int
var a int
var b int
var ans = 1001001001
func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func dfs(pos int) {
	if pos == n {
		ans = int(math.Min(float64(ans), math.Max(float64(a), float64(b))))
		return
	}

	// Aのグループの場合
	a += k[pos]
	dfs(pos+1)
	a -= k[pos]

	// Bのグループの場合
	b += k[pos]
	dfs(pos+1)
	b -= k[pos]
}

func main() {
	n = ScanI()
	k = ScanIArrayWithBlank(n)

	dfs(0)
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