package main

import (
	"flag"

	"github.com/grantmiiller/piazza/pkg/server"
)

func main() {
	var port string
	var query string
	var verbose bool

	flag.StringVar(&port, "p", "8080", "Specify a listening port")
	flag.StringVar(&query, "q", "q", "Specify a query parameter key to decode")
	flag.BoolVar(&verbose, "v", false, "Go loud")

	flag.Parse()
	server.StartServer(port, query, verbose)
}
