package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}

		host, port := line, "443"
		if h, p, err := net.SplitHostPort(line); err == nil {
			host, port = h, p
		}

		c, err := tls.Dial("tcp", fmt.Sprintf("%v:%v", host, port), nil)
		if err != nil {
			log.Fatalf("dial: %v", err)
		}
		if err := c.Handshake(); err != nil {
			log.Fatalf("handshake: %v", err)
		}
		if err := c.Close(); err != nil {
			log.Fatalf("close: %v", err)
		}

		for _, chain := range c.ConnectionState().VerifiedChains {
			for _, cert := range chain {
				if cert.DNSNames != nil {
					fmt.Printf("%-24v %v\n", host, cert.NotAfter)
				}
			}
		}
	}
}
