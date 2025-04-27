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
	S := strings.Split(ScanS(), "")
	count := 0
	var ans []string
	for i := 0; i < len(S); i++ {
		if i == 0 {
			continue
		}
		if S[i] == "|" {
			ans = append(ans, strconv.Itoa(count))
			count = 0
		} else if S[i] == "-" {
			count++
		}
	}

	fmt.Println(strings.Join(ans, " "))
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

func ScanIArray(n int) []int {
	arrI := make([]int, n)
	for i := 0; i < n; i++ {
		arrI[i] = ScanI()
	}

	return arrI
}