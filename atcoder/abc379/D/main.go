package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/liyue201/gostl/ds/queue"
)

var sc = bufio.NewScanner(os.Stdin)
func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	q := ScanI()
	queryList := make([][]int, q)
	for i := 0; i < q; i++ {
		queryList[i] = make([]int, 2)
		queryList[i][0] = ScanI()
		if queryList[i][0] == 1 {
			continue
		}
		queryList[i][1] = ScanI()
	}

	plant := queue.New[int]()
	day := 0

	for _, query := range queryList {
		if query[0] == 1 {
			plant.Push(day)
		} else if query[0] == 2 {
			day += query[1]
		} else if query[0] == 3 {
			num := 0
			for !plant.Empty() {
				if query[1] <= day-plant.Front() {
					plant.Pop()
					num++
				} else {
					break
				}
			}
			// plantが空の時
			fmt.Println(num)
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