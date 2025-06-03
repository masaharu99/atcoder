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
	s := ScanIArrayNotBlank()

	// 1をの番号を管理
	var oneList []int
	for i := 0; i < n; i++ {
		if s[i] == 1 {
			oneList = append(oneList, i+1)
		}
	}

	// 1が1つのみ
	if len(oneList) == 1 {
		fmt.Println(0)
		return
	}

	mid := len(oneList) / 2
	ans := 0

	now := oneList[mid]
	for i := 1; 0 <= mid-i; i++ {
		// fmt.Printf("left, %d, %d, %d\n", now, oneList[mid-i], ans)
		ans += now - oneList[mid-i] - 1
		now--
	}

	now = oneList[mid]
	for i := 1; mid+i < len(oneList); i++ {
		// fmt.Printf("right, %d, %d, %d\n", now, oneList[mid-i], ans)
		ans += oneList[mid+i] - now - 1
		now++
	}

	fmt.Println(ans)
}

func f(oneList []int, mid int) int {
	ans := 0
	cl := oneList[mid]
	cr := oneList[mid]

	for i := 1; ; i++ {
		if 0 <= mid-i && mid-i < len(oneList) {
			// fmt.Println(-1, cl, oneList[mid-i])
			ans += cl - oneList[mid-i] - 1
			cl--
		}
		if 0 <= mid+i && mid+i < len(oneList) {
			// fmt.Println(1, cr, oneList[mid+i])
			ans += oneList[mid+i] - cr - 1
			cr++
		}
		if mid-i < 0 || len(oneList) <= mid-i || mid+i < 0 || len(oneList) <= mid+i {
			break
		}
	}

	return ans
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
