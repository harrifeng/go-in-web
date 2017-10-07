package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/harrifeng/go-in-web/01-bbs/data"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/signup", signup)
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

func signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
