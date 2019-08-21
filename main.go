package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isUpsideDownUint(n uint64) bool {
	input := strconv.FormatUint(n, 10)

	if strings.ContainsAny(input, "23457") {
		return false
	}

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

	str := ""

	for x := len(input) - 1; x >= 0; x-- {
		//fmt.Printf("--x-->%d\n", x)
		rev := m[input[x]]
		if rev == "" {
			return false
		}
		str = str + rev
		//fmt.Println(rev)
	}

	/*
		///fmt.Printf("input:%v   str:%v", input, str)
		if input == str {
			fmt.Printf("input:%v   str:%v\n", input, str)
			//fmt.Println("yep")
		} else {
			//fmt.Println("nope")

		}*/

	return input == str
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
