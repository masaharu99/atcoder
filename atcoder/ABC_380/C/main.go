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
	n, k := ScanI(), ScanI()
	s := strings.Split(ScanS(), "")

	var list [][]int
	num := 0
	for i := 0; i < n; i++ {
		if s[i] == "0" {
			continue
		}

		if i == 0 || s[i-1] == "0" {
			list = append(list, []int{i, 0})
		}
		if i == n-1 || s[i+1] == "0" {
			list[num][1] = i
			num++
		}
	}

	insert := list[k-2][1] + 1
	ks := list[k-1][0]
	kl := list[k-1][1] + 1
	fmt.Printf("%s%s%s%s\n", strings.Join(s[:insert], ""), strings.Join(s[ks:kl], ""), strings.Join(s[insert:ks], ""), strings.Join(s[kl:], ""))
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