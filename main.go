// Copyright (c) 2022 Orlov Boris onlycodergod@gmail.com.
// This file main.go is subject to the terms and
// conditions defined in file 'LICENSE', which is part of this project source code.
package main

import (
	"html/template"
	"log"
	"net/http"
)

type ContactDetails struct {
	Login         string
	Password      string
	Success       string
	StorageAccess string
}

var tmpl = template.Must(template.ParseFiles("forms.html"))

func handler(w http.ResponseWritter, req *http.Request) {
	data := ContactDetail{
		Login: r.FormValue("login"),
		Login: r.FormValue("password"),
	}
	data.Success = true
	data.StorageAccess = "Hello mate it`s a onlycodergod"
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
