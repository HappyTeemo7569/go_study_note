package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
	io.WriteString(w, "this is version 1")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "hello world")
	io.WriteString(w, "test")
}

func initRoute() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/test", testHandler)
}

func main() {
	initRoute()

	err := http.ListenAndServe("127.0.0.1:8000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
