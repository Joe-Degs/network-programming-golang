package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	var host string
	flag.StringVar(&host, "host", "localhost", "hostname to resolve")
	flag.Parse()

	addrs, err := net.LookupHost(host)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(addrs)
}
