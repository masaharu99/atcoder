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
	N := ScanI()
	P := ScanIArrayWithBlank(N)
	Q := ScanIArrayWithBlank(N)

	type humanType struct {
		bib int
		who int
	}
	// 番号で検索
	human := make(map[int]humanType)
	// ゼッケンで番号を検索
	bib := make(map[int]int)
	for i := 0; i < N; i++ {
		human[i+1] = humanType{bib: Q[i], who: P[i]}
		bib[Q[i]] = i + 1
	}

	// Qをソート（ゼッケン）
	sort.Slice(Q, func(i, j int) bool {
		return Q[i] < Q[j]
	})

	for _, val := range Q {
		// 何番の人か
		h := bib[val]
		// 誰を見ているか
		who := human[h].who
		fmt.Printf("%d ", human[who].bib)
	}
	fmt.Println()

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
