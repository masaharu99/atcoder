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
	s := make([][]string, n+1)
	for i := 0; i < len(s); i++ {
		s[i] = make([]string, n+1)
	}

	for i := 0; i < n+1; i++ {
		j := n + 1 - i

		if j < i || i == 0 {
			continue
		}

		fill(s, i, j)
	}

	for i := 0; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 || j == 0 {
				continue
			}
			fmt.Print(s[i][j])
		}
		fmt.Println()
	}
}

func fill(s [][]string, i, j int) {
	c := ""
	if i%2 == 0 {
		c = "."
	} else {
		c = "#"
	}

	for ni := i; ni <= j; ni++ {
		for nj := i; nj <= j; nj++ {
			s[ni][nj] = c
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
