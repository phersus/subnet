package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	v       = []int{2, 4, 8, 16, 32, 64, 128}
	netList = []string{
		"%s.%s.%s.%s/32",
		"%s.%s.%s.%s/31",
		"%s.%s.%s.%s/30",
		"%s.%s.%s.%s/29",
		"%s.%s.%s.%s/28",
		"%s.%s.%s.%s/27",
		"%s.%s.%s.%s/26",
		"%s.%s.%s.%s/25",
		"%s.%s.%s.0/24",
		"%s.%s.%s.0/23",
		"%s.%s.%s.0/22",
		"%s.%s.%s.0/21",
		"%s.%s.%s.0/20",
		"%s.%s.%s.0/19",
		"%s.%s.%s.0/18",
		"%s.%s.%s.0/17",
		"%s.%s.0.0/16",
		"%s.%s.0.0/15",
		"%s.%s.0.0/14",
		"%s.%s.0.0/13",
		"%s.%s.0.0/12",
		"%s.%s.0.0/11",
		"%s.%s.0.0/10",
		"%s.%s.0.0/9",
		"%s.0.0.0/8",
		"%s.0.0.0/7",
		"%s.0.0.0/6",
		"%s.0.0.0/5",
		"%s.0.0.0/4",
		"%s.0.0.0/3",
		"%s.0.0.0/2",
		"%s.0.0.0/1",
	}
)

func main() {

	ip := "172.17.34.80"          // Our initial IPv4 address to analyze
	sOct := make([]string, 4)     // We prepare the placeholder for the individual octets
	sOct = strings.Split(ip, ".") // We separate the octets
	sByte := toInt(sOct)          // We get an int version of the IPv4

	/* ####################### We start computing ############################### */

	// from /32 to /25
	n4 := checkOdd(sByte[3]) // We check the last octet for oddity
	_r4 := netBorders(n4)    // We get all the class C subnets 4th octet
	e := []int{sByte[3]}     // we add sByte[3] to a slice to be able to append to it
	r4 := append(e, _r4...)  // We append the rest of the netBorders to e
	fmt.Println(r4)

	// from /24 to /17
	n3 := checkOdd(sByte[2])
	_r3 := netBorders(n3)
	e = []int{sByte[2]}
	r3 := append(e, _r3...)
	fmt.Println(r3)

	// from /16 to /9
	n2 := checkOdd(sByte[1])
	_r2 := netBorders(n2)
	e = []int{sByte[1]}
	r2 := append(e, _r2...)
	fmt.Println(r2)

	// from /8 to /1
	n1 := checkOdd(sByte[0])
	_r1 := netBorders(n1)
	e = []int{sByte[0]}
	r1 := append(e, _r1...)
	fmt.Println(r1)

	/* Using the function that consolidates these computations */
	n, nStr := crtNetBorders(sByte)
	fmt.Println(n)
	fmt.Println(nStr)

	fmt.Println(crtSubNets(sOct, nStr))

	res := make([]string, 32)

	for i := 0; i <= 7; i++ {
		res[i] = fmt.Sprintf(netList[i], sOct[0], sOct[1], sOct[2], nStr[0][i])
	}

	for i := 8; i <= 15; i++ {
		res[i] = fmt.Sprintf(netList[i], sOct[0], sOct[1], nStr[1][i-8])
	}

	for i := 16; i <= 23; i++ {
		res[i] = fmt.Sprintf(netList[i], sOct[0], nStr[2][i-16])
	}

	for i := 24; i <= 31; i++ {
		res[i] = fmt.Sprintf(netList[i], nStr[3][i-24])
	}
	fmt.Println(res)

	println()
	println()
	println()
	fmt.Println(fmt.Sprintf(netList[0], sOct[0], sOct[1], sOct[2], nStr[0][0]))
	fmt.Println(fmt.Sprintf(netList[1], sOct[0], sOct[1], sOct[2], nStr[0][1]))
	fmt.Println(fmt.Sprintf(netList[2], sOct[0], sOct[1], sOct[2], nStr[0][2]))
	fmt.Println(fmt.Sprintf(netList[3], sOct[0], sOct[1], sOct[2], nStr[0][3]))
	fmt.Println(fmt.Sprintf(netList[4], sOct[0], sOct[1], sOct[2], nStr[0][4]))
	fmt.Println(fmt.Sprintf(netList[5], sOct[0], sOct[1], sOct[2], nStr[0][5]))
	fmt.Println(fmt.Sprintf(netList[6], sOct[0], sOct[1], sOct[2], nStr[0][6]))
	fmt.Println(fmt.Sprintf(netList[7], sOct[0], sOct[1], sOct[2], nStr[0][7]))

	fmt.Println(fmt.Sprintf(netList[8], sOct[0], sOct[1], nStr[1][0]))
	fmt.Println(fmt.Sprintf(netList[9], sOct[0], sOct[1], nStr[1][1]))
	fmt.Println(fmt.Sprintf(netList[10], sOct[0], sOct[1], nStr[1][2]))
	fmt.Println(fmt.Sprintf(netList[11], sOct[0], sOct[1], nStr[1][3]))
	fmt.Println(fmt.Sprintf(netList[12], sOct[0], sOct[1], nStr[1][4]))
	fmt.Println(fmt.Sprintf(netList[13], sOct[0], sOct[1], nStr[1][5]))
	fmt.Println(fmt.Sprintf(netList[14], sOct[0], sOct[1], nStr[1][6]))
	fmt.Println(fmt.Sprintf(netList[15], sOct[0], sOct[1], nStr[1][7]))

	fmt.Println(fmt.Sprintf(netList[16], sOct[0], nStr[2][0]))
	fmt.Println(fmt.Sprintf(netList[17], sOct[0], nStr[2][1]))
	fmt.Println(fmt.Sprintf(netList[18], sOct[0], nStr[2][2]))
	fmt.Println(fmt.Sprintf(netList[19], sOct[0], nStr[2][3]))
	fmt.Println(fmt.Sprintf(netList[20], sOct[0], nStr[2][4]))
	fmt.Println(fmt.Sprintf(netList[21], sOct[0], nStr[2][5]))
	fmt.Println(fmt.Sprintf(netList[22], sOct[0], nStr[2][6]))
	fmt.Println(fmt.Sprintf(netList[23], sOct[0], nStr[2][7]))

	fmt.Println(fmt.Sprintf(netList[24], nStr[3][0]))
	fmt.Println(fmt.Sprintf(netList[25], nStr[3][1]))
	fmt.Println(fmt.Sprintf(netList[26], nStr[3][2]))
	fmt.Println(fmt.Sprintf(netList[27], nStr[3][3]))
	fmt.Println(fmt.Sprintf(netList[28], nStr[3][4]))
	fmt.Println(fmt.Sprintf(netList[29], nStr[3][5]))
	fmt.Println(fmt.Sprintf(netList[30], nStr[3][6]))
	fmt.Println(fmt.Sprintf(netList[31], nStr[3][7]))

}

func crtSubNets(sOct []string, nStr [][]string) []string {

	res := make([]string, 32)

	for i := 0; i <= 7; i++ {
		res[i] = fmt.Sprintf(netList[i], sOct[0], sOct[1], sOct[2], nStr[0][i])
	}

	for i := 8; i <= 15; i++ {
		res[i] = fmt.Sprintf(netList[i], sOct[0], sOct[1], nStr[1][i-8])
	}

	for i := 16; i <= 23; i++ {
		res[i] = fmt.Sprintf(netList[i], sOct[0], nStr[2][i-16])
	}

	for i := 24; i <= 31; i++ {
		res[i] = fmt.Sprintf(netList[i], nStr[3][i-24])
	}
	return res
}

// crtNetBorders creates all the network subnets from /1 to /32
func crtNetBorders(bite []int) ([][]int, [][]string) {
	n := make([]int, 4)
	_r := make([][]int, 4)
	r := make([][]int, 4)

	_rStr := make([][]string, 4)
	rStr := make([][]string, 4)
	biteStr := make([]string, 4)
	biteStr = toStr(bite)

	for i := range r {
		r[i] = make([]int, 8)
		rStr[i] = make([]string, 8)
		for j := range n {
			n[j] = checkOdd(bite[j])
			_r[j] = netBorders(n[j])
			_rStr[j] = toStr(_r[j])
			r[i] = append([]int{bite[3-i]}, _r[3-i]...)
			rStr[i] = append([]string{biteStr[3-i]}, _rStr[3-i]...)
		}
	}
	return r, rStr
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

func toStr(sByte []int) []string {
	sOct := make([]string, 8)
	for i := range sByte {
		sOct[i] = strconv.Itoa(sByte[i])
	}
	return sOct
}

// netBorders will provide a []int with the subnet limits on the given octet
func netBorders(n int) []int {
	r := make([]int, 7)
	for i := range v {
		r[i] = n - n%v[i]
	}
	return r
}
