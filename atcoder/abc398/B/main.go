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
	a := ScanIArrayWithBlank(7)
	arr := make([]int, 14)

	for _, val := range a {
		arr[val]++
	}

	two := 0
	three := 0
	for _, val := range arr {
		if 3 <= val {
			three++
			continue
		}
		if 2 <= val {
			two++
		}
	}

	if (1 <= two && 1 <= three) || (2 <= three) {
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
