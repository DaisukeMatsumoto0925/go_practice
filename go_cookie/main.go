package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
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

var store = sessions.NewCookieStore([]byte("SESSION_KEY"))

func MyHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["foo"] = "bar"
	session.Values[42] = 43

	session.Options.HttpOnly = true
	session.Options.MaxAge = 10000
	session.Options.Secure = true

	session.AddFlash("foo")
	session.AddFlash("bar")

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// http.HandleFunc("/", setCookies)
	// http.HandleFunc("/cookie", showCookie)
	http.HandleFunc("/gorilla", MyHandler)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
