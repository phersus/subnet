package main

import (
	"fmt"
	"subnet/subnetservice"
)

func main() {

	ip := "172.17.34.80"
	fmt.Println(subnetservice.SubNets(ip))

}
