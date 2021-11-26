package main

import (
	"fmt"

	"github.com/phersus/subnet"
)

func main() {

	ip := "57.16.10.24"
	fmt.Println(subnet.CrtSubNets(ip))
}
