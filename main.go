package main

import (
	"fmt"
	"gowebserver/api"
	"gowebserver/data"
	"html/template"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to index"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About me"))
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("views/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	html.Execute(w, data.GetAll())
}

func main() {
	server := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public"))

	server.HandleFunc("/about", aboutHandler)
	server.HandleFunc("/template", templateHandler)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/create", api.Post)
	server.Handle("/", fs)

	err := http.ListenAndServe(":3000", server)

	if err != nil {
		fmt.Println("Error starting server")
	}
}
