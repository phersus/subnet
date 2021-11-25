package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	v       = []int{2, 4, 8, 16, 32, 64, 128}
	netList = []string{
		"%s.0.0.0./1",
		"%s.0.0.0./2",
		"%s.0.0.0./3",
		"%s.0.0.0./4",
		"%s.0.0.0./5",
		"%s.0.0.0./6",
		"%s.0.0.0./7",
		"%s.0.0.0./8",
		"%s.%s.0.0./9",
		"%s.%s.0.0./10",
		"%s.%s.0.0./11",
		"%s.%s.0.0./12",
		"%s.%s.0.0./13",
		"%s.%s.0.0./14",
		"%s.%s.0.0./15",
		"%s.%s.0.0./16",
		"%s.%s.%s.0./17",
		"%s.%s.%s.0./18",
		"%s.%s.%s.0./19",
		"%s.%s.%s.0./20",
		"%s.%s.%s.0./21",
		"%s.%s.%s.0./22",
		"%s.%s.%s.0./23",
		"%s.%s.%s.0./24",
		"%s.%s.%s.%s/25",
		"%s.%s.%s.%s/26",
		"%s.%s.%s.%s/27",
		"%s.%s.%s.%s/28",
		"%s.%s.%s.%s/29",
		"%s.%s.%s.%s/30",
		"%s.%s.%s.%s/31",
		"%s.%s.%s.%s/32",
	}
)

func main() {

	ip := "172.17.34.79"          // Our initial IPv4 address to analyze
	sOct := make([]string, 4)     // We prepare the placeholder for the individual octets
	sOct = strings.Split(ip, ".") // We separate the octets
	sByte := toInt(sOct)          // We get an int version of the IPv4

	/* We start computing */

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

	fmt.Println(fmt.Sprintf(netList[0], nStr[0][7]))

	// for i := range nStr {
	//
	// 	for j := range nStr[i] {
	// 		switch i {
	// 		case 0:
	// 			switch j {
	// 			case 0:
	// 				fmt.Printf(nStr[i][j] + ".0.0.0/8\n")
	// 			case 1:
	// 				fmt.Printf(nStr[i][j] + ".0.0.0/7\n")
	// 			case 2:
	// 				fmt.Printf(nStr[i][j] + ".0.0.0/6\n")
	// 			case 3:
	// 				fmt.Printf(nStr[i][j] + ".0.0.0/5\n")
	// 			case 4:
	// 				fmt.Printf(nStr[i][j] + ".0.0.0/4\n")
	// 			case 5:
	// 				fmt.Printf(nStr[i][j] + ".0.0.0/3\n")
	// 			case 6:
	// 				fmt.Printf(nStr[i][j] + ".0.0.0/2\n")
	// 			case 7:
	// 				fmt.Printf(nStr[i][j] + ".0.0.0/1\n")
	// 			}
	// 		case 1:
	// 			switch j {
	// 			case 0:
	// 				fmt.Printf(sOct[0] + "." + nStr[i][j] + ".0.0/16\n")
	// 			case 1:
	// 				fmt.Printf(sOct[0] + "." + nStr[i][j] + ".0.0/15\n")
	// 			case 2:
	// 				fmt.Printf(sOct[0] + "." + nStr[i][j] + ".0.0/14\n")
	// 			case 3:
	// 				fmt.Printf(sOct[0] + "." + nStr[i][j] + ".0.0/13\n")
	// 			case 4:
	// 				fmt.Printf(sOct[0] + "." + nStr[i][j] + ".0.0/12\n")
	// 			case 5:
	// 				fmt.Printf(sOct[0] + "." + nStr[i][j] + ".0.0/11\n")
	// 			case 6:
	// 				fmt.Printf(sOct[0] + "." + nStr[i][j] + ".0.0/10\n")
	// 			case 7:
	// 				fmt.Printf(sOct[0] + "." + nStr[i][j] + ".0.0/9\n")
	// 			}
	// 		case 2:
	// 			switch j {
	// 			case 0:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + nStr[i][j] + ".0/24\n")
	// 			case 1:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + nStr[i][j] + ".0/23\n")
	// 			case 2:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + nStr[i][j] + ".0/22\n")
	// 			case 3:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + nStr[i][j] + ".0/21\n")
	// 			case 4:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + nStr[i][j] + ".0/20\n")
	// 			case 5:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + nStr[i][j] + ".0/19\n")
	// 			case 6:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + nStr[i][j] + ".0/18\n")
	// 			case 7:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + nStr[i][j] + ".0/17\n")
	// 			}
	// 		case 3:
	// 			switch j {
	// 			case 0:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + sOct[2] + "." + nStr[i][j] + "/32\n")
	// 			case 1:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + sOct[2] + "." + nStr[i][j] + "/31\n")
	// 			case 2:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + sOct[2] + "." + nStr[i][j] + "/30\n")
	// 			case 3:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + sOct[2] + "." + nStr[i][j] + "/29\n")
	// 			case 4:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + sOct[2] + "." + nStr[i][j] + "/28\n")
	// 			case 5:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + sOct[2] + "." + nStr[i][j] + "/27\n")
	// 			case 6:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + sOct[2] + "." + nStr[i][j] + "/26\n")
	// 			case 7:
	// 				fmt.Printf(sOct[0] + "." + sOct[1] + "." + sOct[2] + "." + nStr[i][j] + "/25\n")
	// 			}
	// 		}
	// 	}
	// }

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
		// for i := len(r) - 1; i >= 0; i-- {
		r[i] = make([]int, 8)
		rStr[i] = make([]string, 8)

		for j := range n {
			// for j := len(n) - 1; j >= 0; j-- {
			n[j] = checkOdd(bite[j])
			_r[j] = netBorders(n[j])
			_rStr[j] = toStr(_r[j])
			r[i] = append([]int{bite[3-i]}, _r[3-i]...)
			// r[i] = append([]int{bite[i]}, _r[i]...)
			rStr[i] = append([]string{biteStr[3-i]}, _rStr[3-i]...)
			// rStr[i] = append([]string{biteStr[i]}, _rStr[i]...)
			// fmt.Println(r[i])
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
