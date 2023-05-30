package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	filePath := "ip_addresses.txt"

	// Open the input file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ipStr := scanner.Text()
		ip := net.ParseIP(ipStr)

		if ip == nil {
			fmt.Printf("Invalid IP address: %s\n", ipStr)
			continue
		}

		ipRange := getIPRange(ip)
		fmt.Printf("IP: %s, Range: %s - %s\n", ipStr, ipRange.Start, ipRange.End)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read file: %s\n", err)
		return
	}
}

// IPRange represents the range of IP addresses
type IPRange struct {
	Start net.IP
	End   net.IP
}

// getIPRange returns the range of IP addresses for a given IP
func getIPRange(ip net.IP) IPRange {
	mask := ip.DefaultMask()
	startIP := ip.Mask(mask)
	endIP := make(net.IP, len(startIP))
	copy(endIP, startIP)

	for i := range endIP {
		endIP[i] |= ^mask[i]
	}

	return IPRange{
		Start: startIP,
		End:   endIP,
	}
}
