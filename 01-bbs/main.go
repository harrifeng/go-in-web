package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/harrifeng/go-in-web/01-bbs/data"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	server := &http.Server{
		Addr:    "0.0.0.0:7890",
		Handler: mux,
	}
	server.ListenAndServe()

	fmt.Println()
	os.Exit(0)
}

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()

	fmt.Println(threads, err)
	fmt.Fprintf(writer, "Hello World, %s", request.URL.Path[1:])
}
