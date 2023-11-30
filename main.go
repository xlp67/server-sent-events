package main

import (
"html/template"
"log"
"net/http"
"sse/sse"
)



func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/partial-response", func(w http.ResponseWriter, r *http.Request){
		sse.ResponseSSE("text", w, r)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		tpl, err := template.ParseFiles("index.html")
		if err != nil {log.Fatal(err)}
		tpl.Execute(w, nil)
	})
	http.ListenAndServe(":8080", mux)


}	