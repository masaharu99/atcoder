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
	s := make([][]string, 8)
	for i := 0; i < 8; i++ {
		s[i] = strings.Split(ScanS(), "")
	}

	exitI := make([]bool, 8)
	exitJ := make([]bool, 8)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if s[i][j] == "#" {
				exitI[i] = true
				exitJ[j] = true
			}
		}
	}

	ans := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if exitI[i] == false && exitJ[j] == false {
				ans++
			}
		}
	}

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