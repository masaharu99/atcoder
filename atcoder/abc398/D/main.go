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
	_ = ScanI()
	r := ScanI()
	c := ScanI()
	s := strings.Split(ScanS(), "")

	dx, dy := 0, 0
	fire := []int{0, 0}
	takahashi := []int{r, c}
	stomes := set.New(comparator.StringComparator, set.WithGoroutineSafe())

	for _, v := range s {
		if v == "N" {
			fire[0]++
		}
		if v == "S" {
			fire[0]--
		}
		if v == "W" {
			fire[1]++
		}
		if v == "E" {
			fire[1]--
		}

		stomes.Insert(addR + "/" + addC)

	}

	fmt.Println()
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
