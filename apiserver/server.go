package apiserver

import (
	"io"
	"log"
	"net/http"

	"github.com/chlins/obs/pkg/register"
)

// Start api server
func Start(l string) {
	http.HandleFunc("/objects/", handle)
	register.Prepare()
	log.Fatal(http.ListenAndServe(l, nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	// select one data server
	dataserver, err := register.RandomSelectDataServer()
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	r.URL.Host = dataserver
	client := new(http.Client)
	req, err := http.NewRequest(r.Method, "http:"+r.URL.String(), r.Body)
	if err == nil {
		resp, err := client.Do(req)
		if err == nil {
			io.Copy(w, resp.Body)
			defer resp.Body.Close()
			return
		}
	}

	w.WriteHeader(http.StatusInternalServerError)
}
