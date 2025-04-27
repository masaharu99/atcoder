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
	n := strings.Split(ScanS(), "")
	var c1, c2, c3 int

	for _, num := range n {
		if num == "1" {
			c1++
		} else if num == "2" {
			c2++
		} else if num == "3" {
			c3++
		}
	}

	if c1 == 1 && c2 == 2 && c3 == 3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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

func ScanIArray(n int) []int {
	arrI := make([]int, n)
	for i := 0; i < n; i++ {
		arrI[i] = ScanI()
	}

	return arrI
}