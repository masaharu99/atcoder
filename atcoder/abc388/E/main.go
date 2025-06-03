package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	a := ScanIArrayWithBlank(n)

	ans := 0
	pushed := make(map[int]bool)

	for i := 0; i < len(a); i++ {
		if _,ok := pushed[i]; ok {
			continue
		}

		result := sort.Search(len(a), func(j int) bool {
			return a[j] - a[i]*2 >= 0
		})
		if result != len(a) {
			if _, ok := pushed[result]; !ok {
				pushed[result] = true
				ans++
			} else {
				for j := result+1; j < len(a); j++ {
					if _, ok := pushed[j]; !ok {
						pushed[j] = true
						ans++
						break
					}
				}
			}
		} else {
			break
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