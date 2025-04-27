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
	m := ScanI()
	x := ScanIArrayWithBlank(m)
	a := ScanIArrayWithBlank(m)

	type Stones struct {
		x, a int
	}
	stones := make([]Stones, m)
	for i := 0; i < m; i++ {
		stones[i] = Stones{x[i], a[i]}
	}

	sort.Slice(stones, func(i, j int) bool {
		return stones[i].x < stones[j].x
	})
	// n+1に石を追加
	stones = append(stones, Stones{n+1, 1})

	ans := 0
	px := 0
	num := 1

	for i := 0; i < len(stones); i++ {
		l := stones[i].x - px
		carry := num - l
		ans += (l-1)*l/2
		ans += carry*l

		if carry < 0 {
			fmt.Println(-1)
			return
		}

		px = stones[i].x
		num = stones[i].a + carry
	}

	if num != 1 {
		fmt.Println(-1)
		return
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