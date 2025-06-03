package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/liyue201/gostl/ds/stack"
)

var sc = bufio.NewScanner(os.Stdin)
var s string
var count int

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	s = ScanS()
	stack := stack.New[rune]()
	ans := "Yes"

	for _, v := range s {
		if v == '(' || v == '[' || v == '<' {
			stack.Push(v)
		} else {
			if stack.Empty() {
				ans = "No"
				break
			}
			prev := stack.Pop()
			if !judge(prev, v) {
				ans = "No"
				break
			}
		}
	}

	if !stack.Empty() {
		ans = "No"
	}

	fmt.Println(ans)
}

func judge(left, right rune) bool {
	if left == '(' && right == ')' {
		return true
	}
	if left == '[' && right == ']' {
		return true
	}
	if left == '<' && right == '>' {
		return true
	}

	return false
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
