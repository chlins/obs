package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var storagePath string

func init() {
	storagePath = os.Getenv("STORAGE_ROOT") + "/objects/"
}

// Handler handles http requests
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		// put method store object
		put(w, r)
	case http.MethodGet:
		// get method get object
		get(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func put(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create(storagePath + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("store error: %s\n", err.Error())
		return
	}
	log.Printf("[debug] store obj (%s)\n", f.Name())
	defer f.Close()
	io.Copy(f, r.Body)
}

func get(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(storagePath + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get error: %s\n", err.Error())
		return
	}
	log.Printf("[debug] get obj (%s)\n", f.Name())
	defer f.Close()
	io.Copy(w, f)
}
