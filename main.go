package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	v    = []int{2, 4, 8, 16, 32, 64, 128}
	list = map[string]int{
		`/31`: 2,
		`/30`: 4,
		`/29`: 8,
		`/28`: 16,
		`/27`: 32,
		`/26`: 64,
		`/25`: 128,
	}
)

func main() {

	ip := "172.17.34.79"          // Our initial IPv4 address to analyze
	sOct := make([]string, 4)     // We prepare the placeholder for the individual octets
	sOct = strings.Split(ip, ".") // We separate the octets

	sByte := toInt(sOct) // We get an int version of the IPv4

	n4 := checkOdd(sByte[3]) // We check the last octet for oddity
	_r4 := netBorders(n4)    // We get all the class C subnets 4th octet
	e := []int{sByte[3]}     //
	r4 := append(e, _r4...)
	fmt.Println(r4)

	n3 := checkOdd(sByte[2])
	_r3 := netBorders(n3)
	e = []int{sByte[2]}
	r3 := append(e, _r3...)
	fmt.Println(r3)

	n2 := checkOdd(sByte[1])
	_r2 := netBorders(n2)
	e = []int{sByte[1]}
	r2 := append(e, _r2...)
	fmt.Println(r2)

	n1 := checkOdd(sByte[0])
	_r1 := netBorders(n1)
	e = []int{sByte[0]}
	r1 := append(e, _r1...)
	fmt.Println(r1)
}

// checkOdd will receive an integer and returns it if it is even, otherwise it returns the
// given integer-1
func checkOdd(i int) int {
	n := i
	if n%2 != 0 {
		return n - 1
	}
	return n
}

// toInt provides the integer version of the IPv4 given as a string
func toInt(sOct []string) []int {
	sByte := make([]int, 4)
	for i := range sOct {
		sByte[i], _ = strconv.Atoi(sOct[i])
	}
	return sByte
}

func netBorders(n int) []int {
	r := make([]int, 7)
	for i := range v {
		r[i] = n - n%v[i]
	}
	return r
}
