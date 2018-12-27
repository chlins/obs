package main

import (
	"flag"

	"github.com/chlins/obs/dataserver"
)

var l = flag.String("l", ":7878", "listen port")

func main() {
	flag.Parse()

	dataserver.Start(*l)
}
