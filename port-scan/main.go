package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

// Scan checks if a port is open for given host and protocol
func Scan(protocol, host string, port int) bool {
	conn, err := net.DialTimeout(protocol, host+":"+strconv.Itoa(port), 500*time.Millisecond)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	var wg sync.WaitGroup
	for port := 0; port <= 65535; port++ {
		wg.Add(1)
		go func(protocol, host string, port int, wg *sync.WaitGroup) {
			isOpen := Scan(protocol, host, port)
			if isOpen {
				fmt.Printf("port: %d is open", port)
			}
			wg.Done()
		}("tcp", "localhost", port, &wg)
	}
	wg.Wait()
}
