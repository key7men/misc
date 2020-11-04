package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("starting server...")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(writer, `Hello, you want know how to debugger go with docker?`)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}