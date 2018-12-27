package dataserver

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/chlins/obs/pkg/objects"
	"github.com/chlins/obs/pkg/register"
)

// Start a data server
func Start(l string) {
	http.HandleFunc("/objects/", objects.Handler)
	log.Printf("server listen in %s\n", l)
	err := register.DataServer(getLocalIP() + l)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(l, nil))
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var ip string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	return ip
}
