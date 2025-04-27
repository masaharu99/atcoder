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
	s := strings.Split(ScanS(), "")

	max := 0

	for i := 0; i < n; i++ {
		if s[i] != "/" {
			continue
		}

		d := 1
		for {
			if i-d < 0 || n <= i+d {
				break
			}
			if s[i-d] != "1" || s[i+d] != "2" {
				break
			}
			d++
		}
		num := 1 + 2 * (d - 1)
		if max < num {
			max = num
		}
	}

	fmt.Println(max)
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