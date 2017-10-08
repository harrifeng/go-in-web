package main

import (
	"fmt"
	"html/template"
	"log"
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
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/authenticate", authenticate)
	server := &http.Server{
		Addr:    "0.0.0.0:7890",
		Handler: mux,
	}
	server.ListenAndServe()

	fmt.Println()
	os.Exit(0)
}

func index(writer http.ResponseWriter, request *http.Request) {

	fmt.Println(data.UserByEmail("hfeng@hfeng.com"))

	threads, err := data.Threads()

	fmt.Println(threads)
	if err != nil {
		fmt.Println(writer, request, "Cannot get threads")
	} else {
		generateHTML(writer, threads, "layout", "public.navbar", "index")
	}

}

func login(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "login")
}

func authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		fmt.Println("authenticate:", err)
	} else {
		http.Redirect(writer, request, "/", 302)
	}
	if user.Password == request.PostFormValue("password") {

	} else {
		http.Redirect(writer, request, "/", 302)
	}
}

func signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

func signupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}

	if err := user.Create(); err != nil {
		log.Fatal(err)
	}

	http.Redirect(writer, request, "/index", 302)

}

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
