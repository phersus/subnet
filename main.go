package main

import (
	"fmt"
	"strconv"
	"strings"
)

var v = []int{2, 4, 8, 16, 32, 64, 128}

func main() {
	ipAdd := "172.18.52.111"
	nets, err := GenNets(ipAdd)
	if err != nil {
		return
	}
	fmt.Println(nets)

}

// Octets takes an IPv4 address as a string and breaks it into
// 4 octets returned in a []int
func Octets(ip string) ([]int, error) {
	r := make([]int, 0, 24)
	oct := strings.Split(ip, ".")
	fmt.Println(oct)
	for k, v := range oct {
		fmt.Printf("value: %s Type: %T", v, k)
		println()
	}
	return r, nil
}

func GenNets(ip string) ([]string, error) {
	r := make([]string, 8)
	r[0] = ip + "/32"
	sOctet, err := Octets(ip)
	// sOctet := make([]int, len(sByte))
	// for i := range sByte {
	// 	sOctet[i] = int(sByte[i])
	// }

	if err != nil {
		return []string{}, err
	}
	for i := 1; i <= 7; i++ {
		switch i {
		case 1:
			filling("/31", r, sOctet, 1)
		case 2:
			filling("/30", r, sOctet, 2)
		case 3:
			filling("/29", r, sOctet, 3)
		case 4:
			filling("/28", r, sOctet, 4)
		case 5:
			filling("/27", r, sOctet, 5)
		case 6:
			filling("/26", r, sOctet, 6)
		case 7:
			filling("/25", r, sOctet, 7)

		}

	}
	return r, nil
}

func filling(a string, r []string, sOctet []int, j int) {

	if sOctet[3]%v[j-1] == 0 {
		r[j] = strconv.Itoa(sOctet[3]) + a
		return
	}
	r[j] = strconv.Itoa(sOctet[3]-(sOctet[3]%v[j-1])) + a
}
