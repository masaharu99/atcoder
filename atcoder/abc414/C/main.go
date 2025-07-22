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
	a := ScanI()
	n := ScanI()

	fmt.Println(solve(n, a))
}

func createPalindromicList(limit int) []string {
	vl := make([]string, 0)

	// 1,2桁
	for a := 1; a <= 9; a++ {
		add1 := strconv.Itoa(a)
		if limit >= a {
			vl = append(vl, add1)
		}

		add2 := add1 + add1
		v, _ := strconv.Atoi(add2)
		if limit >= v {
			vl = append(vl, add2)
		}
	}

	// 3,4桁
	for a := 1; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			three := strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(a)
			v, _ := strconv.Atoi(three)
			if limit >= v {
				vl = append(vl, three)
			}

			four := strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(b) + strconv.Itoa(a)
			v, _ = strconv.Atoi(four)
			if limit >= v {
				vl = append(vl, four)
			}
		}
	}

	// 5,6桁
	for a := 1; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				as := strconv.Itoa(a)
				bs := strconv.Itoa(b)
				cs := strconv.Itoa(c)

				v, _ := strconv.Atoi(as + bs + cs + bs + as)
				if limit >= v {
					vl = append(vl, as+bs+cs+bs+as)
				}

				v2, _ := strconv.Atoi(as + bs + cs + cs + bs + as)
				if limit >= v2 {
					vl = append(vl, as+bs+cs+cs+bs+as)
				}
			}
		}
	}

	// 7,8桁
	for a := 1; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					as := strconv.Itoa(a)
					bs := strconv.Itoa(b)
					cs := strconv.Itoa(c)
					ds := strconv.Itoa(d)

					v, _ := strconv.Atoi(as + bs + cs + ds + cs + bs + as)
					if limit >= v {
						vl = append(vl, as+bs+cs+ds+cs+bs+as)
					}

					v2, _ := strconv.Atoi(as + bs + cs + ds + ds + cs + bs + as)
					if limit >= v2 {
						vl = append(vl, as+bs+cs+ds+ds+cs+bs+as)
					}
				}
			}
		}
	}

	// 9,10桁
	for a := 1; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					for e := 0; e <= 9; e++ {
						as := strconv.Itoa(a)
						bs := strconv.Itoa(b)
						cs := strconv.Itoa(c)
						ds := strconv.Itoa(d)
						es := strconv.Itoa(e)

						v, _ := strconv.Atoi(as + bs + cs + ds + es + ds + cs + bs + as)
						if limit >= v {
							vl = append(vl, as+bs+cs+ds+es+ds+cs+bs+as)
						}

						v2, _ := strconv.Atoi(as + bs + cs + ds + es + es + ds + cs + bs + as)
						if limit >= v2 {
							vl = append(vl, as+bs+cs+ds+es+es+ds+cs+bs+as)
						}
					}
				}
			}
		}
	}

	// 11,12桁
	for a := 1; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					for e := 0; e <= 9; e++ {
						for f := 0; f <= 9; f++ {
							as := strconv.Itoa(a)
							bs := strconv.Itoa(b)
							cs := strconv.Itoa(c)
							ds := strconv.Itoa(d)
							es := strconv.Itoa(e)
							fs := strconv.Itoa(f)

							v, _ := strconv.Atoi(as + bs + cs + ds + es + fs + es + ds + cs + bs + as)
							if limit >= v {
								vl = append(vl, as+bs+cs+ds+es+fs+es+ds+cs+bs+as)
							}

							v2, _ := strconv.Atoi(as + bs + cs + ds + es + fs + fs + es + ds + cs + bs + as)
							if limit >= v2 {
								vl = append(vl, as+bs+cs+ds+es+fs+fs+es+ds+cs+bs+as)
							}
						}
					}
				}
			}
		}
	}

	return vl
}

func solve(limit, n int) int {
	ans := 0

	for a := 1; a <= 9; a++ {
		as := strconv.Itoa(a)
		v1 := a
		v2, _ := strconv.Atoi(as + as)
		if limit >= v1 && isPalindromic(as, n) {
			ans += v1
		}
		if limit >= v2 && isPalindromic(as+as, n) {
			ans += v2
		}
		for b := 0; b <= 9; b++ {
			bs := strconv.Itoa(b)
			c1 := as + bs + as
			v1, _ := strconv.Atoi(c1)
			c2 := as + bs + bs + as
			v2, _ := strconv.Atoi(c2)
			if limit >= v1 && isPalindromic(c1, n) {
				ans += v1
			}
			if limit >= v2 && isPalindromic(c2, n) {
				ans += v2
			}
			for c := 0; c <= 9; c++ {
				cs := strconv.Itoa(c)
				c1 := as + bs + cs + bs + as
				v1, _ := strconv.Atoi(c1)
				c2 := as + bs + cs + cs + bs + as
				v2, _ := strconv.Atoi(c2)
				if limit >= v1 && isPalindromic(c1, n) {
					ans += v1
				}
				if limit >= v2 && isPalindromic(c2, n) {
					ans += v2
				}
				for d := 0; d <= 9; d++ {
					ds := strconv.Itoa(d)
					c1 := as + bs + cs + ds + cs + bs + as
					v1, _ := strconv.Atoi(c1)
					c2 := as + bs + cs + ds + ds + cs + bs + as
					v2, _ := strconv.Atoi(c2)
					if limit >= v1 && isPalindromic(c1, n) {
						ans += v1
					}
					if limit >= v2 && isPalindromic(c2, n) {
						ans += v2
					}
					for e := 0; e <= 9; e++ {
						es := strconv.Itoa(e)
						c1 := as + bs + cs + ds + es + ds + cs + bs + as
						v1, _ := strconv.Atoi(c1)
						c2 := as + bs + cs + ds + es + es + ds + cs + bs + as
						v2, _ := strconv.Atoi(c2)
						if limit >= v1 && isPalindromic(c1, n) {
							ans += v1
						}
						if limit >= v2 && isPalindromic(c2, n) {
							ans += v2
						}
						for f := 0; f <= 9; f++ {
							fs := strconv.Itoa(f)
							c1 := as + bs + cs + ds + es + fs + es + ds + cs + bs + as
							v1, _ := strconv.Atoi(c1)
							c2 := as + bs + cs + ds + es + fs + fs + es + ds + cs + bs + as
							v2, _ := strconv.Atoi(c2)
							if limit >= v1 && isPalindromic(c1, n) {
								ans += v1
							}
							if limit >= v2 && isPalindromic(c2, n) {
								ans += v2
							}
						}
					}
				}
			}
		}
	}

	return ans
}

func isPalindromic(tenNumberStr string, n int) bool {
	str := convertNNumber(tenNumberStr, n)

	if str == ReverseStr(str) {
		return true
	}

	return false
}

func convertNNumber(bStr string, n int) string {
	b, _ := strconv.Atoi(bStr)
	str := ""

	for {
		str += strconv.Itoa(b % n)
		b = b / n
		if b == 0 {
			break
		}
	}

	return str
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

type UnionFind struct {
	n    int
	root []int
}

func NewUnionFind(n int) UnionFind {
	root := make([]int, n+1)
	for i := 0; i < n; i++ {
		root[i] = -1
	}

	return UnionFind{n, root}
}

func (uf UnionFind) Leader(x int) int {
	if uf.root[x] < 0 {
		return x
	}

	uf.root[x] = uf.Leader(uf.root[x])
	return uf.root[x]
}

func (uf UnionFind) Merge(x, y int) int {
	xr := uf.Leader(x)
	yr := uf.Leader(y)
	if xr == yr {
		return xr
	}

	if -uf.root[xr] < -uf.root[yr] {
		xr, yr = yr, xr
	}
	uf.root[xr] += uf.root[yr]
	uf.root[yr] = xr

	return xr
}

func (uf UnionFind) Same(x, y int) bool {
	if uf.Leader(x) == uf.Leader(y) {
		return true
	}

	return false
}

func ReverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
