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
	a := ScanIArrayWithBlank(n)

	num1 := judge(1, n, a)
	num2 := judge(2, n, a)
	if num1 <= num2 {
		fmt.Println(num2)
	} else {
		fmt.Println(num1)
	}
}

func judge(start, n int, a []int) int {
	count := 0
	s := set.New(comparator.IntComparator, set.WithGoroutineSafe())
	ans := set.New(comparator.IntComparator, set.WithGoroutineSafe())
	ans.Insert(0)

	for i := start; i < n; i += 2 {
		if len(a) <= i {
			break
		}
		if a[i-1] == a[i] {
			if s.Contains(a[i]) {
				s.Clear()
				ans.Insert(count)
				s.Insert(a[i])
				count = 2
			} else {
				s.Insert(a[i])
				count += 2
			}
		} else {
			s.Clear()
			ans.Insert(count)
			count = 0
		}

		if count != 0 {
			ans.Insert(count)
		}
	}

	return ans.Last().Value()
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
