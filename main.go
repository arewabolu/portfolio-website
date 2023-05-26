package main

import (
	"net/http"
	"text/template"
)

type info struct {
	Name   string
	Email  string
	Github string
}

func DefaultHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp := template.Must(template.ParseFiles("./frontend/index.html"))
		data := info{
			Name:   "Arewa Bolu",
			Email:  "bolu.arewa@gmail.com",
			Github: "arewabolu",
		}
		err := temp.Execute(w, data)
		if err != nil {
			w.Write([]byte("unable to complete request"))
		}
	}

}

func main() {
	http.HandleFunc("/", DefaultHandle)
	styles := http.FileServer(http.Dir("./frontend/style"))
	http.Handle("/style/", http.StripPrefix("/style/", styles))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
