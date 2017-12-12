package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var conc = flag.Int("workers", 8, "number of fetches to perform concurrently")

func main() {
	flag.Parse()
	work := make(chan job)
	go func() {
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

			work <- job{host, port}
		}
		close(work)
	}()

	wg := sync.WaitGroup{}
	sema := make(chan bool, *conc)
	for w := range work {
		sema <- true
		wg.Add(1)
		go func(j job) {
			defer func() {
				wg.Done()
				<-sema
			}()

			res, err := getDate(j.host, j.port)
			if err != nil {
				log.Printf("get date: %v", err)
				return
			}
			fmt.Printf("%-24v %v\n", res.host, res.exp)
		}(w)
	}
	wg.Wait()
}

type job struct {
	host string
	port string
}

type res struct {
	host string
	exp  time.Time
}

func getDate(host, port string) (res, error) {
	r := res{
		host: fmt.Sprintf("%v:%v", host, port),
	}
	c, err := tls.Dial("tcp", r.host, nil)
	if err != nil {
		return r, errors.Wrap(err, "dial")
	}
	if err := c.Handshake(); err != nil {
		return r, errors.Wrap(err, "handshake")
	}
	if err := c.Close(); err != nil {
		return r, errors.Wrap(err, "close")
	}

	for _, chain := range c.ConnectionState().VerifiedChains {
		for _, cert := range chain {
			if cert.DNSNames != nil {
				r.exp = cert.NotAfter
			}
		}
	}
	return r, nil
}
