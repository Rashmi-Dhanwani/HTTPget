package main

import (
	"io"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "rashmi")
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/POST/myname"] = hello

	server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	if strings.Compare(r.URL.String(), "/POST/myname") == 0 {
		io.WriteString(w, "rashmi")
	} else {
		io.WriteString(w, "Please enter proper input")
	}

}
