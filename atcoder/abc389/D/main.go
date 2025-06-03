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
	r := ScanI()

	x := 0
	sum := 0

	for y := r-1; y >= 0; y-- {
		for ; judge(x+1, y, r); x++ {}
		sum += x
	}

	sum *= 4
	sum++
	fmt.Println(sum)
}

func judge(x, y, r int) bool {
	nx := (2*x+1)*(2*x+1)
	ny := (2*y+1)*(2*y+1)
	if nx+ny <= 4*r*r {
		return true
	}
	return false
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