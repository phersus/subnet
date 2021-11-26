package subnetservice

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	v       = []int{2, 4, 8, 16, 32, 64, 128}
	netList = "%s.%s.%s.%s/%d"
	re      = regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
)

// CrtSubNets create all the subnet for a given IPv4 address as a string
// the result is an organized []string from the most specific network (/32)
// at position 0 of the slice
func CrtSubNets(ip string) []string {
	if !re.MatchString(ip) {
		log.Println("ip address not valid!")
		return nil
	}
	// return []string{""}

	sOct := strings.Split(ip, ".") // We separate the octets
	sByte := toInt(sOct)           // We get an int version of the IPv4
	_, nStr := crtNetBorders(sByte)

	res := make([]string, 32)
	for i := 0; i <= 7; i++ {
		res[i] = fmt.Sprintf(netList, sOct[0], sOct[1], sOct[2], nStr[0][i], 32-i)
	}
	for i := 8; i <= 15; i++ {
		res[i] = fmt.Sprintf(netList, sOct[0], sOct[1], nStr[1][i-8], "0", 32-i)
	}
	for i := 16; i <= 23; i++ {
		res[i] = fmt.Sprintf(netList, sOct[0], nStr[2][i-16], "0", "0", 32-i)
	}
	for i := 24; i <= 31; i++ {
		res[i] = fmt.Sprintf(netList, nStr[3][i-24], "0", "0", "0", 32-i)
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
	sByte := make([]int, len(sOct))
	for i := range sOct {
		sByte[i], _ = strconv.Atoi(sOct[i]) // No need to check error, fair checked done on the regex
	}
	return sByte
}

func toStr(sByte []int) []string {
	sOct := make([]string, len(sByte))
	for i := range sByte {
		sOct[i] = strconv.Itoa(sByte[i])
	}
	return sOct
}

// netBorders will provide a []int with the subnet limits on the given octet
func netBorders(n int) []int {
	r := make([]int, len(v))
	for i := range v {
		r[i] = n - n%v[i]
	}
	return r
}
