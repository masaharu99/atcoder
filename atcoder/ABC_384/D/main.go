package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/liyue201/gostl/ds/set"
	"github.com/liyue201/gostl/utils/comparator"
)

var sc = bufio.NewScanner(os.Stdin)
func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	n := ScanI()
	S := ScanI()
	a := ScanIArrayWithBlank(n)

	sumList := make([]int, 2*n+1)
	for i := 1; i < 2*n+1; i++ {
		sumList[i] = sumList[i-1] + a[i%n]
	}

	rm := S % sumList[n]

	s := set.New(comparator.IntComparator, set.WithGoroutineSafe())
	for i := 0; i < 2*n+1; i++ {
		s.Insert(sumList[i])
		if s.Contains(sumList[i]-rm) {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
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