package main

import (
	"fmt"
	"subnet/subnetservice"
)

func main() {

	ip := "57.16.10.24"
	fmt.Println(subnetservice.CrtSubNets(ip))
}
