package main

import (
	"fmt"
	"math"
)

// var sc = bufio.NewScanner(os.Stdin)
// func init() {
// 	sc.Buffer([]byte{}, math.MaxInt64)
// 	sc.Split(bufio.ScanWords)
// }

func main_bk() {
	n := ScanI()
	a := ScanIArrayWithBlank(n)

	s := 0
	r := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] += s
		if a[i] == 0 {
			fmt.Print(0, " ")
			continue
		}

		num := int(math.Min(float64(a[i]), float64(n-1-i)))
		a[i] -= num
		s = s + 1 - r[i]
		r[i+num]++

		fmt.Print(a[i], " ")
	}
	fmt.Println()
}

// func ScanI() int {
// 	sc.Scan()
// 	str, _ := strconv.Atoi(sc.Text())
// 	return str
// }

// func ScanS() string {
// 	sc.Scan()
// 	return sc.Text()
// }

// func ScanIArrayNotBlank() []int {
// 	str := strings.Split(ScanS(), "")
// 	arrI := make([]int, len(str))
// 	for i := 0; i < len(str); i++ {
// 		num, _ := strconv.Atoi(str[i])
// 		arrI[i] = num
// 	}

// 	return arrI
// }

// func ScanIArrayWithBlank(n int) []int {
// 	arrI := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		arrI[i] = ScanI()
// 	}

// 	return arrI
// }