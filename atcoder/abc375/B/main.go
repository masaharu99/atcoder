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
	n := ScanI()
	xy := make([][]int, n)
	for i := 0; i < n; i++ {
		xy[i] = ScanIArrayWithBlank(2)
	}
	cx, cy := 0, 0
	var ans float64

	for i := 0; i < n; i++ {
		nxy := xy[i]
		a := (cx-nxy[0]) * (cx-nxy[0])
		b := (cy-nxy[1]) * (cy-nxy[1])
		ans += math.Sqrt(float64(a+b))
		cx, cy = nxy[0], nxy[1]
	}

	a := (cx-0) * (cx-0)
	b := (cy-0) * (cy-0)
	ans += math.Sqrt(float64(a+b))

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