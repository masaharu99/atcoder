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
var a [][]string
var na [][]string
func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	n = ScanI()
	a = make([][]string, n+1)
	na = make([][]string, n+1)
	a[0] = []string{}
	na[0] = []string{}
	for i := 0; i < n; i++ {
		add := []string{"X"}
		add = append(add, strings.Split(ScanS(), "")...)
		a[i+1] = add
		na[i+1] = make([]string, n+1)
	}

	operate()

	for _, val := range na {
		if len(val) == 0 {
			continue
		}
		fmt.Println(strings.Join(val[1:], ""))
	}
}

func operate() {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			ni, nj := rotation(i, j)
			na[ni][nj] = a[i][j]
		}
	}
}

func rotation(i, j int) (int, int) {
	di := int(math.Min(float64(i), float64(n+1-i)))
	dj := int(math.Min(float64(j), float64(n+1-j)))
	num := int(math.Min(float64(di), float64(dj)))
	ni, nj := i, j
	for k := 0; k < num%4; k++ {
		ni, nj = nj, n + 1 - ni
	}
	return ni, nj
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