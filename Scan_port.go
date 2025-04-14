package main

import (
	"fmt"
	"net"
	"time"
)

func scanPort(host string, port int, results chan int) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		results <- 0
		return
	}
	conn.Close()
	fmt.Printf("Port %d is open\n", port)
	results <- 1

}

func main() {
	host := "scanme.nmap.org"
	openPorts := 0
	results := make(chan int)

	for port := 0; port < 100; port++ {
		go scanPort(host, port, results)
	}

	for i := 0; i < 100; i++ {
		openPorts += <-results
	}

	fmt.Println("==========================")
	fmt.Println("Scan completed")
	fmt.Println("%d ports are open.\n", openPorts)
	fmt.Println("==========================")
}
