package main

import (
	"log"
	"net/http"
	"text/template"
)

func setCookies(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "hoge",
		Value: "bar",
	}
	http.SetCookie(w, cookie)
}

func showCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("hoge")
	if err != nil {
		log.Fatal("cookie:", err)
	}

	tmpl := template.Must(template.ParseFiles("./cookie.html"))
	tmpl.Execute(w, cookie)
}
