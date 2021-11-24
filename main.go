package main

import (
	"fmt"
	"strconv"
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

func filling(a string, r []string, sOctet []int, j int) {

	if sOctet[3]%v[j-1] == 0 {
		r[j] = strconv.Itoa(sOctet[3]) + a
		return
	}
	r[j] = strconv.Itoa(sOctet[3]-(sOctet[3]%v[j-1])) + a
}
