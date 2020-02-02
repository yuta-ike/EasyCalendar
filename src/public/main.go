package main

import (
	"html/template"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/*", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("clock.tpl"))

	m := map[string]string{
		"Date": time.Now().Format("2006-01-02"),
		"Time": time.Now().Format("15:04:05"),
	}

	tpl.Execute(w, m)

	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("<html><h>HELLO</h><p>world</p></html>"))
}
