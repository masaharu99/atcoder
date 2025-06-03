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
	s := strings.Split(ScanS(), "")

	if len(s) % 2 == 0 {
		fmt.Println("No")
		return
	}

	ans := "Yes"
	mid := (len(s) + 1) / 2 - 1
	for i := 0; i < n; i++ {
		if i < mid && s[i] != "1" {
			ans = "No"
		}
		if i == mid && s[i] != "/" {
			ans = "No"
		}
		if mid < i && s[i] != "2" {
			ans = "No"
		}
	}

	fmt.Println(ans)
}

func Scans() {
	panic("unimplemented")
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

func ScanIArray(n int) []int {
	arrI := make([]int, n)
	for i := 0; i < n; i++ {
		arrI[i] = ScanI()
	}

	return arrI
}