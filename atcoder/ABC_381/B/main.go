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
	s := strings.Split(ScanS(), "")

	ans := "Yes"
	if len(s) % 2 == 1 {
		ans = "No"
	}

	kv := make(map[string]int)

	for i := 0; i < len(s)-1; i = i+2 {
		if s[i] != s[i+1] {
			ans = "No"
		} else {
			key := s[i]
			_, ok := kv[key]
			if ok {
				ans = "No"
			} else {
				kv[key] = 0
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

func ScanIArray(n int) []int {
	arrI := make([]int, n)
	for i := 0; i < n; i++ {
		arrI[i] = ScanI()
	}

	return arrI
}