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
	k := ScanI()
	s := strings.Split(ScanS(), "")

	type Rle struct {
		state string
		num int
	}
	var rle []Rle
	ele := -1

	for i := 0; i < n; i++ {
		if i != 0 && s[i] == s[i-1] {
			rle[ele].num++
		} else {
			ele++
			rle = append(rle, Rle{s[i], 1})
		}
	}

	ans := 0
	for _, val := range rle {
		if val.state == "O" && k <= val.num {
			ans += val.num / k
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

func ScanIArray(n int) []int {
	arrI := make([]int, n)
	for i := 0; i < n; i++ {
		arrI[i] = ScanI()
	}

	return arrI
}