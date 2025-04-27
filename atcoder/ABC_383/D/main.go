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
	N := ScanI()
	sieve := eratos(int(2e6+1))
	ans := 0

	for _, p := range sieve {
		for _, q := range sieve {
			if p <= q {
				break
			}
			if p*p*q*q > N {
				break
			}
			ans++
		}
	}

	for _, p := range sieve {
		if p*p*p*p*p*p*p*p > N {
			break
		}
		ans++
	}

	fmt.Println(ans)
}

// 確認したい数+1を引数指定する
// 戻り値に1は含まない
func eratos(n int) []int {
	// 素数判定のリスト
	isPrime := make([]bool, n)
	for i := 1; i < n; i++ {
		isPrime[i] = true
	}

	for i := 2; i < n; i++ {
		if isPrime[i] {
			for j := i*2; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	var sieve []int
	for i := 1; i < n; i++ {
		if isPrime[i] && i != 1 {
			sieve = append(sieve, i)
		}
	}

	return sieve
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