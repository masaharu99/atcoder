package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/liyue201/gostl/ds/queue"
	"github.com/liyue201/gostl/ds/stack"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	S, T := ScanS(), ScanS()
	// X := make([]string, 0, len(S))

	q := queue.New[int]()
	s := stack.New[int]()
	for i := 0; i < len(S); i++ {
		if S[i] == T[i] {
			continue
		}
		if S[i] > T[i] {
			q.Push(i)
		} else {
			s.Push(i)
		}
	}
	fmt.Println(q.Size() + s.Size())

	cur := S
	for !q.Empty() {
		i := q.Pop()
		cur = cur[:i] + string(T[i]) + cur[i+1:]
		fmt.Println(cur)
		// X = append(X, cur)
	}
	for !s.Empty() {
		i := s.Pop()
		cur = cur[:i] + string(T[i]) + cur[i+1:]
		fmt.Println(cur)
		// X = append(X, cur)
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
