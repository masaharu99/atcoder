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
	h := ScanI()
	w := ScanI()
	s := make([][]string, h)
	for i := 0; i < h; i++ {
		s[i] = strings.Split(ScanS(), "")
	}

	minX := 100000
	maxX := 0
	minY := 100000
	maxY := 0

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] != "#" {
				continue
			}

			if minX > j {
				minX = j
			}
			if maxX < j {
				maxX = j
			}
			if minY > i {
				minY = i
			}
			if maxY < i {
				maxY = i
			}
		}
	}

	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			if s[i][j] == "." {
				fmt.Println("No")
				return
			}
		}
	}

	fmt.Println("Yes")
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