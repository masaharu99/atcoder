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
	s := ScanS()
	lcnt := make([]int, 26)
	rcnt := make([]int, 26)

	for _, val := range s {
		rcnt[val-'A']++
	}

	ans := 0

	for i := 0; i < len(s); i++ {
		rcnt[s[i]-'A']--
		for j := 0; j < 26; j++ {
			ans += lcnt[j] * rcnt[j]
		}
		lcnt[s[i]-'A']++
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
