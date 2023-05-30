package main

import (
	"fmt"
	"net"
	"sync"
)

func resolveDomain(domain string, wg *sync.WaitGroup, results chan string) {
	defer wg.Done()

	ipAddr, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		results <- fmt.Sprintf("Failed to resolve %s: %s", domain, err)
		return
	}

	results <- fmt.Sprintf("%s: %s", domain, ipAddr.IP)
}

func main() {
	domains := []string{"example.com", "google.com", "github.com"}

	results := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(domains))

	for _, domain := range domains {
		go resolveDomain(domain, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
