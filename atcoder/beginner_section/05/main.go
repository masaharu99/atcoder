package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n,a,b int
	fmt.Scan(&n,&a,&b)

	var ans int
	for i:=1; i<=n; i++ {
		str := strconv.Itoa(i)
		if len(str) == 1 {
			if a<=i && i<=b {
				ans += i
			}
		} else {
			arr := strings.Split(str, "")
			var sum int
			for j:=0; j<len(arr); j++ {
				iInt, _ := strconv.Atoi(arr[j])
				sum += iInt
			}

			if a<=sum && sum<=b {
				ans += i
			}
		}
	}

	fmt.Println(ans)
}