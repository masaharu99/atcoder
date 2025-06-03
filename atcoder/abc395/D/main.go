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
var (
	n           int
	q           int
	op          [][]int
	pegionToBox []int
	boxToNest   []int
	nestToBox   []int
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	n = ScanI()
	q = ScanI()
	op = make([][]int, q)
	for i := 0; i < len(op); i++ {
		opn := ScanI()
		if opn == 3 {
			a := ScanI()
			op[i] = []int{opn, a}
		} else {
			a := ScanI()
			b := ScanI()
			op[i] = []int{opn, a, b}
		}
	}

	pegionToBox = make([]int, n+1)
	nestToBox = make([]int, n+1)
	boxToNest = make([]int, n+1)
	for i := 0; i < n+1; i++ {
		pegionToBox[i] = i
		nestToBox[i] = i
		boxToNest[i] = i
	}

	for _, query := range op {
		if query[0] == 1 {
			q1(query[1], query[2])
		} else if query[0] == 2 {
			q2(query[1], query[2])
		} else {
			q3(query[1])
		}
	}
}

func q1(a, b int) {
	pegionToBox[a] = boxToNest[b]
}

func q2(a, b int) {
	fmt.Println("bofore")
	fmt.Println("boxToNest", nestToBox)
	fmt.Println("nestToBox", boxToNest)
	nestToBox[a], nestToBox[b] = nestToBox[b], nestToBox[a]
	boxToNest[nestToBox[a]], boxToNest[nestToBox[b]] = boxToNest[nestToBox[b]], boxToNest[nestToBox[a]]
	fmt.Println("after")
	fmt.Println("boxToNest", nestToBox)
	fmt.Println("nestToBox", boxToNest)
}

func q3(a int) {
	current := pegionToBox[a]
	fmt.Println(boxToNest[current])
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
