package main

import (
	"fmt"
	utils_ip "go-demo/utils/ip"
)

func main() {
	ip := utils_ip.GetInternetIP()
	fmt.Println(ip)
	fmt.Println(utils_ip.GetLocalIP())
}
