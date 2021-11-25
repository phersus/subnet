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

// CrtSubNets create all the subnet for a given IPv4 address as a string
// the result is an organized []string from the most specific network (/32)
// at position 0 of the slice
func CrtSubNets(ip string) []string {

	re := regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	if !re.MatchString(ip) || len(ip) == 0 {
		log.Println("ip address not valid!")
		return []string{""}
	}

	sOct := make([]string, 4)     // We prepare the placeholder for the individual octets
	sOct = strings.Split(ip, ".") // We separate the octets
	sByte := toInt(sOct)          // We get an int version of the IPv4
	_, nStr := crtNetBorders(sByte)

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
