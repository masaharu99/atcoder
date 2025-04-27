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
var n int
var k [][]int

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	n = ScanI()
	k = make([][]int, n)
	for i := 0; i < n; i++ {
		men := ScanI()
		k[i] = ScanIArrayWithBlank(men)
	}
	// ランレングス圧縮
	ran := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		eleMap := make(map[int]int)
		for _, val := range k[i] {
			if _, exist := eleMap[val]; !exist {
				eleMap[val] = 1
			} else {
				eleMap[val]++
			}
		}
		ran[i] = eleMap
	}

	ans := 0.0

	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			// サイコロが同じ場合
			if i == j {
				continue
			}
			mom := len(k[i]) * len(k[j])
			san := 0

			keys := make([]int, 0, len(ran[i]))
			for k := range ran[i] {
				keys = append(keys, k)
			}
			// iのキーでjを検索
			for _, key := range keys {
				if _, exist := ran[j][key]; exist {
					san += ran[i][key] * ran[j][key]
				}
			}
			probability := float64(san) / float64(mom)
			ans = math.Max(ans, probability)
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
