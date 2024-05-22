package main

import (
	"fmt"
	"net"
)

func main() {

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Iterate over each network interface
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Print the interface name and associated IP addresses
		fmt.Println("Interface:", iface.Name)
		for _, addr := range addrs {
			// Check if the address is an IP address
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			// Skip loopback and link-local addresses
			// if ipNet.IP.IsLoopback() || ipNet.IP.IsLinkLocalUnicast() {
			// 	continue
			// }

			// Print the non-loopback, non-link-local IP address
			fmt.Println("  IP Address:", ipNet.IP)
		}
	}
}
