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

type Snake struct {
	head int
	length int
}

func main() {
	q := ScanI()
	query := make([][]int, q)
	for i := 0; i < q; i++ {
		first := ScanI()
		if first != 2 {
			second := ScanI()
			query[i] = []int{first, second}
		} else {
			query[i] = []int{first}
		}
	}

	var snakeList []Snake
	start := 0
	change := 0

	for _,val := range query {
		fmt.Println(snakeList)
		if val[0] == 1 {
			if len(snakeList) == 0 {
				snakeList = append(snakeList, Snake{head: 0, length: val[1]})
			} else {
				lastSnake := snakeList[len(snakeList)-1]
				head := lastSnake.head + lastSnake.length
				snakeList = append(snakeList, Snake{head: head, length: val[1]})
			}
		} else if val[0] == 2 {
			if len(snakeList) >= 2 {
				change += snakeList[start].length
			} else {
				change = 0
			}
			start++
		} else {
			tg := val[1] - 1 + start
			fmt.Println(snakeList[tg].head - change)
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