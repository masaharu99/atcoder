package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	N := ScanI()
	M := ScanI()
	human := ScanIArray(N)
	susi := ScanIArray(M)

	for i := 0; i < N-1; i++ {
		human[i+1] = int(math.Min(float64(human[i]), float64(human[i+1])))
	}

	for _, s := range susi {
		index := sort.Search(len(human), func(i int) bool {
			return human[i] <= s
		})

		if index == len(human) {
			fmt.Println(-1)
		} else {
			fmt.Println(index+1)
		}
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