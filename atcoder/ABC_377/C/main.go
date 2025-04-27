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
	m := ScanI()
	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = ScanI()
		b[i] = ScanI()
	}

	kv := make(map[string]bool)

	type Directions struct {
		i, j int
	}
	directions := []Directions{
		{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1},
	}

	for i := 0; i < m; i++ {
		key := strconv.Itoa(a[i]) + "/" + strconv.Itoa(b[i])
		kv[key] = true
		for _, direction := range directions {
			ki := a[i]+direction.i
			kj := b[i]+direction.j
			if 0<ki && ki<=n && 0<kj && kj<=n {
				key := strconv.Itoa(ki) + "/" + strconv.Itoa(kj)
				kv[key] = true
			}
		}
	}

	fmt.Println(n*n - len(kv))
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