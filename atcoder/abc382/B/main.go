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
	n := ScanI()
	d := ScanI()
	s := strings.Split(ScanS(), "")

	for day := 0; day < d; day++ {
		for i := n -1; 0 <= i; i-- {
			if s[i] == "@" {
				s[i] = "."
				break
			}
		}
	}

	fmt.Println(strings.Join(s, ""))
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