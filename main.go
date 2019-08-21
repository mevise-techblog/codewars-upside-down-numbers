package main

import (
	"fmt"
	"strconv"
)

func isUpsideDownUint(n uint64) bool {
	input := strconv.FormatUint(n, 10)

	var m map[byte]string
	m = make(map[byte]string)

	m['0'] = "0"
	m['1'] = "1"
	m['2'] = ""
	m['3'] = ""
	m['4'] = ""
	m['5'] = ""
	m['6'] = "9"
	m['7'] = ""
	m['8'] = "8"
	m['9'] = "6"

	// grab the two outside numbers
	for i := 0; i < len(input)/2; i++ {
		c1 := input[i]
		c2 := input[len(input)-1-i]
		if c1 != '1' && c1 != '6' && c1 != '8' && c1 != '9' {
			return false
		}
		if c2 != '1' && c2 != '6' && c2 != '8' && c2 != '9' {
			return false
		}

		if m[c1] != c2 {
			return false
		}
	}
	return true
}

func UpsideDown(n1, n2 string) uint64 {

	i1, _ := strconv.ParseUint(n1, 10, 64)
	i2, _ := strconv.ParseUint(n2, 10, 64)

	var cnt uint64 = 0

	for x := i1; x <= i2; x++ {
		if isUpsideDownUint(x) {
			cnt++
		}
	}

	return cnt
}

func main() {
	fmt.Println(UpsideDown("100000", "12345678900000000"))
	//fmt.Println(isUpsideDownUint(16891))

}
