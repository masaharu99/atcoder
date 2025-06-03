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
	bn := ScanI()
	wn := ScanI()
	b := make([]int, bn)
	w := make([]int, wn)
	for i := 0; i < bn; i++ {
		b[i] = ScanI()
	}
	for i := 0; i < wn; i++ {
		w[i] = ScanI()
	}

	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
	sort.Slice(w, func(i, j int) bool { return w[i] > w[j] })

	nb := 0
	ans := 0

	for i := 0; i < bn; i++ {
		tg := b[i]
		if 0 <= tg {
			ans += tg
			nb++
		} else {
			break
		}
	}

	for i := 0; i < wn; i++ {
		if wn <= i {
			break
		}

		tg := w[i]
		if 0 <= tg {
			// WよりBの方が選んだ数が多い
			if i < nb {
				ans += tg
				// 次のBの数 < 次のWの数
			} else if nb < bn && 0 < tg+b[nb] {
				ans += tg + b[nb]
				nb++
			} else {
				break
			}
		} else {
			break
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
