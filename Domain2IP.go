package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
	inputFile := "input.txt"
	outputFile := "output.txt"

	// Open input file
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Failed to open input file: %s\n", err)
		return
	}
	defer input.Close()

	// Open output file
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Failed to create output file: %s\n", err)
		return
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)

	domains := []string{}
	for scanner.Scan() {
		domain := scanner.Text()
		domains = append(domains, domain)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read input file: %s\n", err)
		return
	}

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
		fmt.Fprintln(output, result)
	}

	fmt.Printf("Results saved to %s\n", outputFile)
}
