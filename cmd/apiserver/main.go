package main

import (
	"flag"

	"github.com/chlins/obs/apiserver"
)

var l = flag.String("l", ":6868", "listen port")

func main() {
	flag.Parse()

	apiserver.Start(*l)
}
