package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/liyue201/gostl/ds/stack"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	Q := ScanI()
	q := make([][]int, Q)
	for i := 0; i < Q; i++ {
		o := ScanI()
		if o == 1 {
			q[i] = []int{o, ScanI()}
		} else {
			q[i] = []int{o}
		}
	}

	s := stack.New[int]()
	for i := 0; i < 100; i++ {
		s.Push(0)
	}

	for _, query := range q {
		if query[0] == 1 {
			s.Push(query[1])
		} else {
			fmt.Println(s.Pop())
		}
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
