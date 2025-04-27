package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	s1 := set.New[int](comparator.IntComparator, set.WithGoroutineSafe())
	s1.Insert(1)
	s1.Insert(2)
	s2 := set.New[int](comparator.IntComparator, set.WithGoroutineSafe())
	s2.Insert(1)
	s3 := set.New[int](comparator.IntComparator, set.WithGoroutineSafe())
	s3.Insert(1)
	s3.Insert(2)
	s3.Insert(3)
	S := []*set.Set[int]{s1, s2, s3}
	sort.Slice(S, func(i, j int) bool {
		return S[i].Size() < S[j].Size()
	})
	fmt.Println(S)
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
