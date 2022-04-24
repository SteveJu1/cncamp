package httpserver

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func NewServer(addr string) {
	http.HandlerFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
}

func injectRequestHeader(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		w.Header().Set(key, strings.Join(value, ";"))
	}
	w.Write([]byte("request header in response header"))
}

func getEnv(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("VERSION")

	w.Header().Set("Version", version)
	w.Write([]byte("version in response header"))

}
