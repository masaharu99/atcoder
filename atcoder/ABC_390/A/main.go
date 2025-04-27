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
	A := ScanIArrayWithBlank(5)

	for i := 0; i < 4; i++ {
		na := create(i, A)
		prev := na[0]
		result := true
		for i := 1; i < 5; i++ {
			if prev > na[i] {
				result = false
			}
			prev = na[i]
		}
		if result {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}

func create(i int, A []int) []int {
	var na []int
	na = append(na, A[:i]...)
	na = append(na, A[i+1])
	na = append(na, A[i])
	na = append(na, A[i+2:]...)

	return na
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