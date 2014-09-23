package main

import (
	"io"
	"log"
	"net/http"
	"server/ser"
	"time"
)

func main() {

	s := &http.Server{
		Addr: ":9393",
		// Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/test", ser.Hello)
	log.Fatal(s.ListenAndServe())
}

func myHandler(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "test hello")
}
