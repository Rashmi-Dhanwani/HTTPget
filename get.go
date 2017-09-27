package main

import (
	"io"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/GET/myname?name="] = hello

	server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	stringSlice := strings.Split(r.URL.String(), "=")

	if len(stringSlice) > 1 {
		io.WriteString(w, stringSlice[1])
	} else {
		io.WriteString(w, "Please enter proper input")
	}

}
