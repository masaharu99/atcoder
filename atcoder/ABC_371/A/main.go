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
	sab := ScanS()
	sac := ScanS()
	sbc := ScanS()

	if sab == "<" && sbc == "<" {
		fmt.Println("B")
	}
	if sac == "<" && sbc == ">" {
		fmt.Println("C")
	}
	if sab == ">" && sac == "<" {
		fmt.Println("A")
	}
	if sbc == "<" && sac == ">" {
		fmt.Println("C")
	}
	if sac == ">" && sab == "<" {
		fmt.Println("A")
	}
	if sbc == ">" && sab == ">" {
		fmt.Println("B")
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
