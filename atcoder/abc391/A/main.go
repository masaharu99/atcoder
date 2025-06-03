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
	d := ScanS()
	ans := ""

	if d == "N" {
		ans = "S"
	} else if d == "S" {
		ans = "N"
	} else if d == "E" {
		ans = "W"
	} else if d == "W" {
		ans = "E"
	} else if d == "NE" {
		ans = "SW"
	} else if d == "SW" {
		ans = "NE"
	} else if d == "NW" {
		ans = "SE"
	} else if d == "SE" {
		ans = "NW"
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